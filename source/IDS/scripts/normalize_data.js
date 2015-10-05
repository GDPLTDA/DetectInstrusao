var atk ={ack : "dos",buffer_overflow : "u2r",ftp_write : "r2l",guess_passwd : "r2l",imap : "r2l",ipsweep : "probe",land : "dos",loadmodule : "u2r",multihop : "r2l",neptune : "dos",nmap : "probe",perl : "u2r",phf : "r2l",pod : "dos",portsweep : "probe",rootkit : "u2r",satan : "probe",smurf : "dos",spy : "r2l",teardrop : "dos",warezclient : "r2l",warezmaster : "r2l"};
var simb = {    "protocol_type" : [         "tcp",         "udp",         "icmp"    ],    "service" : [         "http",         "smtp",         "domain_u",         "auth",         "finger",         "telnet",         "eco_i",         "ftp",         "ntp_u",         "ecr_i",         "other",         "urp_i",         "private",         "pop_3",         "ftp_data",         "netstat",         "daytime",         "ssh",         "echo",         "time",         "name",         "whois",         "domain",         "mtp",         "gopher",         "remote_job",         "rje",         "ctf",         "supdup",         "link",         "systat",         "discard",         "X11",         "shell",         "login",         "imap4",         "nntp",         "uucp",         "pm_dump",         "IRC",         "Z39_50",         "netbios_dgm",         "ldap",         "sunrpc",         "courier",         "exec",         "bgp",         "csnet_ns",         "http_443",         "klogin",         "printer",         "netbios_ssn",         "pop_2",         "nnsp",         "efs",         "hostnames",         "uucp_path",         "sql_net",         "vmnet",         "iso_tsap",         "netbios_ns",         "kshell",         "urh_i",         "http_2784",         "harvest",         "aol",         "tftp_u",         "http_8001",         "tim_i",         "red_i"    ],    "flag" : [         "SF",         "S2",         "S1",         "S3",         "OTH",         "REJ",         "RSTO",         "S0",         "RSTR",         "RSTOS0",         "SH"    ],    "land" : [         "0",         "1"    ],    "logged_in" : [         "1",         "0"    ],    "is_host_login" : [         "0",         "1"    ],    "is_guest_login" : [         "0",         "1"    ]};
var val={    "duration" : {        "min" : 0.0000000000000000,        "max" : 58329.0000000000000000    },    "src_bytes" : {        "min" : 0.0000000000000000,        "max" : 1379963888.0000000000000000    },    "dst_bytes" : {        "min" : 0.0000000000000000,        "max" : 1309937401.0000000000000000    },    "wrong_fragment" : {        "min" : 0.0000000000000000,        "max" : 3.0000000000000000    },    "urgent" : {        "min" : 0.0000000000000000,        "max" : 14.0000000000000000    },    "hot" : {        "min" : 0.0000000000000000,        "max" : 77.0000000000000000    },    "num_failed_logins" : {        "min" : 0.0000000000000000,        "max" : 5.0000000000000000    },    "num_compromised" : {        "min" : 0.0000000000000000,        "max" : 7479.0000000000000000    },    "root_shell" : {        "min" : 0.0000000000000000,        "max" : 1.0000000000000000    },    "su_attempted" : {        "min" : 0.0000000000000000,        "max" : 2.0000000000000000    },    "num_root" : {        "min" : 0.0000000000000000,        "max" : 7468.0000000000000000    },    "num_file_creations" : {        "min" : 0.0000000000000000,        "max" : 43.0000000000000000    },    "num_shells" : {        "min" : 0.0000000000000000,        "max" : 2.0000000000000000    },    "num_access_files" : {        "min" : 0.0000000000000000,        "max" : 9.0000000000000000    },    "num_outbound_cmds" : {        "min" : 0.0000000000000000,        "max" : 0.0000000000000000    },    "count" : {        "min" : 0.0000000000000000,        "max" : 511.0000000000000000    },    "srv_count" : {        "min" : 0.0000000000000000,        "max" : 511.0000000000000000    },    "serror_rate" : {        "min" : 0.0000000000000000,        "max" : 0.0000000000000000    },    "srv_serror_rate" : {        "min" : 0.0000000000000000,        "max" : 0.0000000000000000    },    "rerror_rate" : {        "min" : 0.0000000000000000,        "max" : 0.0000000000000000    },    "srv_rerror_rate" : {        "min" : 0.0000000000000000,        "max" : 0.0000000000000000    },    "same_srv_rate" : {        "min" : 0.0000000000000000,        "max" : 0.0000000000000000    },    "diff_srv_rate" : {        "min" : 0.0000000000000000,        "max" : 0.0000000000000000    },    "srv_diff_host_rate" : {        "min" : 0.0000000000000000,        "max" : 0.0000000000000000    },    "dst_host_count" : {        "min" : 0.0000000000000000,        "max" : 255.0000000000000000    },    "dst_host_srv_count" : {        "min" : 0.0000000000000000,        "max" : 255.0000000000000000    },    "dst_host_same_srv_rate" : {        "min" : 0.0000000000000000,        "max" : 0.0000000000000000    },    "dst_host_diff_srv_rate" : {        "min" : 0.0000000000000000,        "max" : 0.0000000000000000    },    "dst_host_same_src_port_rate" : {        "min" : 0.0000000000000000,        "max" : 0.0000000000000000    },    "dst_host_srv_diff_host_rate" : {        "min" : 0.0000000000000000,        "max" : 0.0000000000000000    },    "dst_host_serror_rate" : {        "min" : 0.0000000000000000,        "max" : 0.0000000000000000    },    "dst_host_srv_serror_rate" : {        "min" : 0.0000000000000000,        "max" : 0.0000000000000000    },    "dst_host_rerror_rate" : {        "min" : 0.0000000000000000,        "max" : 0.0000000000000000    },    "dst_host_srv_rerror_rate" : {        "min" : 0.0000000000000000,        "max" : 0.0000000000000000    }};
var tatk = ["normal","dos","probe","r2l","u2r"];


function totaliza(e){
    for (x in e){
      	if (x!=="_id"){
      		var data = new Object();

      		if (simb[x]!==undefined)
      			{
      				data[x] = simb[e[x]] * (1/simb.lenght);
      			}
      		else if (val[x]!==undefined)
      			{
      				data[x] =  e[x]/val[x].max;	

      			}
      		else if (x=="attack")
      		{
      			var atkName = atk[e[x]];
      			data[x] = {
      				"normal" : atkName==="normal" ? 1 : 0,
      				"dos" : atkName==="dos" ? 1 : 0,
      				"probe" : atkName==="probe" ? 1 : 0,
      				"r2l" : atkName==="r2l" ? 1 : 0,
      				"u2r" : atkName==="u2r" ? 1 : 0
      			};	
      		}

      	}
    }

    db.normalizedData.insert(data);

}

db.KDDCup.find().forEach(totaliza);
