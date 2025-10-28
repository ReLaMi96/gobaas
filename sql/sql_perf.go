package sql

import (
	"fmt"
	"time"

	"github.com/ReLaMi96/gobaas/components"
	"github.com/ReLaMi96/gobaas/utils"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	"gorm.io/gorm"
)

func QueryPerfRead(db gorm.DB) ([]components.QueryPerf, error) {
	var result []components.QueryPerf
	err := db.Raw("SELECT query, calls, total_exec_time, mean_exec_time, rows FROM pg_stat_statements ORDER BY queryid DESC LIMIT 100;").Scan(&result).Error
	return result, err
}

func GetDBdetails(db *gorm.DB) (*utils.DBdetails, error) {

	dbname, err := GetDatabaseName(db)
	if err != nil {
		return nil, err
	}
	dbversion, err := GetDatabaseVersion(db)
	if err != nil {
		return nil, err
	}
	host, err := GetDatabaseHost(db)
	if err != nil {
		return nil, err
	}
	port, err := GetDatabasePort(db)
	if err != nil {
		return nil, err
	}
	sslmode, err := GetDatabaseSSLmode(db)
	if err != nil {
		return nil, err
	}

	uptime, err := GetDatabaseUptime(db)
	if err != nil {
		return nil, err
	}

	cpu, err := GetSystemCPU(db)
	if err != nil {
		return nil, err
	}

	ram, err := GetSystemRAM(db)
	if err != nil {
		return nil, err
	}

	space, err := GetSystemDiskSpace(db)
	if err != nil {
		return nil, err
	}

	status, err := CheckDatabaseHealth(db)
	if err != nil {
		return nil, err
	}

	DBdetails := &utils.DBdetails{
		Status:    status,
		DBname:    dbname,
		DBversion: dbversion,
		Host:      host,
		Port:      port,
		SSLmode:   sslmode,
		Uptime:    uptime,
		CPU:       cpu,
		RAM:       ram,
		Space:     space,
	}

	return DBdetails, nil
}

func GetDatabaseName(db *gorm.DB) (string, error) {
	var dbName string
	row := db.Raw("SELECT current_database()").Row()
	if err := row.Scan(&dbName); err != nil {
		return "", err
	}
	return dbName, nil
}

func GetDatabaseHost(db *gorm.DB) (string, error) {
	var host string
	row := db.Raw("SELECT inet_server_addr()").Row()
	if err := row.Scan(&host); err != nil {
		return "", err
	}
	return host, nil
}

func GetDatabasePort(db *gorm.DB) (string, error) {
	var port string
	row := db.Raw("SELECT inet_server_port()").Row()
	if err := row.Scan(&port); err != nil {
		return "", err
	}
	return port, nil
}

func GetDatabaseVersion(db *gorm.DB) (string, error) {
	var version string
	row := db.Raw("SELECT left(current_setting('server_version'),4)").Row()
	if err := row.Scan(&version); err != nil {
		return "", err
	}
	return version, nil
}

func GetDatabaseSSLmode(db *gorm.DB) (string, error) {
	var sslmode string
	row := db.Raw("SHOW ssl").Row()
	if err := row.Scan(&sslmode); err != nil {
		return "", err
	}
	return sslmode, nil
}

func GetDatabaseUptime(db *gorm.DB) (string, error) {
	var uptime string
	row := db.Raw("SELECT ROUND(EXTRACT(epoch FROM (now() - pg_postmaster_start_time())) / 86400, 1) || ' days' as uptime").Row()
	if err := row.Scan(&uptime); err != nil {
		return "", err
	}
	return uptime, nil
}

func GetSystemCPU(db *gorm.DB) (string, error) {
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%.0f%%", cpuPercent[0]), nil
}

func GetSystemRAM(db *gorm.DB) (string, error) {
	virtualMem, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%.0f%%", virtualMem.UsedPercent), nil
}

func GetSystemDiskSpace(db *gorm.DB) (string, error) {
	diskStat, err := disk.Usage("/")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%.0f%%", diskStat.UsedPercent), nil
}

func CheckDatabaseHealth(db *gorm.DB) (string, error) {
	sqlDB, err := db.DB()
	if err != nil {
		return "Disconnected", err
	}

	if err := sqlDB.Ping(); err != nil {
		return "Disconnected", err
	}

	return "Connected", nil
}

func Stats(name string, db *gorm.DB) (string, error) {

	switch name {
	case "uptime":
		return GetDatabaseUptime(db)
	case "cpu":
		return GetSystemCPU(db)
	case "ram":
		return GetSystemRAM(db)
	case "space":
		return GetSystemDiskSpace(db)
	}

	return "", nil
}
