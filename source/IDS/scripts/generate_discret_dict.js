var simb{
"protocol_type" : [],
"service" : [],
"flag" : [],
"land" : [],
"logged_in" : [],
"su_attempted " : [],
"is_host_login" : [],
"is_guest_login" : []
}

function totaliza(e){
    for (x in simb){
        if ( simb[x].indefOf(e[x])>=0 )
        	simb[x].push(e[x]);
    }
}

db.KDDCup.find().forEach(totaliza);



simb



----------------------------------------------

/* 1 */
{
    "protocol_type" : [ 
        "tcp", 
        "udp", 
        "icmp"
    ],
    "service" : [ 
        "http", 
        "smtp", 
        "domain_u", 
        "auth", 
        "finger", 
        "telnet", 
        "eco_i", 
        "ftp", 
        "ntp_u", 
        "ecr_i", 
        "other", 
        "urp_i", 
        "private", 
        "pop_3", 
        "ftp_data", 
        "netstat", 
        "daytime", 
        "ssh", 
        "echo", 
        "time", 
        "name", 
        "whois", 
        "domain", 
        "mtp", 
        "gopher", 
        "remote_job", 
        "rje", 
        "ctf", 
        "supdup", 
        "link", 
        "systat", 
        "discard", 
        "X11", 
        "shell", 
        "login", 
        "imap4", 
        "nntp", 
        "uucp", 
        "pm_dump", 
        "IRC", 
        "Z39_50", 
        "netbios_dgm", 
        "ldap", 
        "sunrpc", 
        "courier", 
        "exec", 
        "bgp", 
        "csnet_ns", 
        "http_443", 
        "klogin", 
        "printer", 
        "netbios_ssn", 
        "pop_2", 
        "nnsp", 
        "efs", 
        "hostnames", 
        "uucp_path", 
        "sql_net", 
        "vmnet", 
        "iso_tsap", 
        "netbios_ns", 
        "kshell", 
        "urh_i", 
        "http_2784", 
        "harvest", 
        "aol", 
        "tftp_u", 
        "http_8001", 
        "tim_i", 
        "red_i"
    ],
    "flag" : [ 
        "SF", 
        "S2", 
        "S1", 
        "S3", 
        "OTH", 
        "REJ", 
        "RSTO", 
        "S0", 
        "RSTR", 
        "RSTOS0", 
        "SH"
    ],
    "land" : [ 
        "0", 
        "1"
    ],
    "logged_in" : [ 
        "1", 
        "0"
    ],
    "is_host_login" : [ 
        "0", 
        "1"
    ],
    "is_guest_login" : [ 
        "0", 
        "1"
    ]
}