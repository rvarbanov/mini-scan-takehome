CREATE TABLE IF NOT EXISTS scan (
    ip VARCHAR(15) NOT NULL,
    port INTEGER NOT NULL,
    service VARCHAR(50) NOT NULL,
    data VARCHAR(1000) NOT NULL,

    timestamp INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT scans_unique_ip_port_service UNIQUE(ip, port, service)
);

CREATE INDEX idx_scan_ip ON scan(ip);
CREATE INDEX idx_scan_service ON scan(service);
CREATE INDEX idx_scan_created_at ON scan(created_at); 
