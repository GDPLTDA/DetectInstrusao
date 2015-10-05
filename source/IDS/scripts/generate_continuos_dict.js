var tots = {
"duration" : {min:null, max:null},
"src_bytes" : {min:null, max:null},
"dst_bytes" : {min:null, max:null},
"wrong_fragment" : {min:null, max:null},
"urgent" : {min:null, max:null},
"hot" : {min:null, max:null},
"num_failed_logins" : {min:null, max:null},
"num_compromised" : {min:null, max:null},
"root_shell" : {min:null, max:null},
"su_attempted" : {min:null, max:null},
"num_root" : {min:null, max:null},
"num_file_creations" : {min:null, max:null},
"num_shells" : {min:null, max:null},
"num_access_files" : {min:null, max:null},
"num_outbound_cmds" : {min:null, max:null},
"count" : {min:null, max:null},
"srv_count" : {min:null, max:null},
"serror_rate" : {min:null, max:null},
"srv_serror_rate" : {min:null, max:null},
"rerror_rate" : {min:null, max:null},
"srv_rerror_rate" : {min:null, max:null},
"same_srv_rate" : {min:null, max:null},
"diff_srv_rate" : {min:null, max:null},
"srv_diff_host_rate" : {min:null, max:null},
"dst_host_count" : {min:null, max:null},
"dst_host_srv_count" : {min:null, max:null},
"dst_host_same_srv_rate" : {min:null, max:null},
"dst_host_diff_srv_rate" : {min:null, max:null},
"dst_host_same_src_port_rate" : {min:null, max:null},
"dst_host_srv_diff_host_rate" : {min:null, max:null},
"dst_host_serror_rate" : {min:null, max:null},
"dst_host_srv_serror_rate" : {min:null, max:null},
"dst_host_rerror_rate" : {min:null, max:null},
"dst_host_srv_rerror_rate" : {min:null, max:null}
};

function totaliza(e){
    for (x in tots){
            if (tots[x].min === null || e[x] < tots[x].min )
                tots[x].min = e[x];
            
            if (tots[x].max === null || e[x] > tots[x].max)
                tots[x].max = e[x];
    }
}

db.KDDCup.find().forEach(totaliza);



tots



--------------------------------------------------------------



/* 1 */
{
    "duration" : {
        "min" : 0.0000000000000000,
        "max" : 58329.0000000000000000
    },
    "src_bytes" : {
        "min" : 0.0000000000000000,
        "max" : 1379963888.0000000000000000
    },
    "dst_bytes" : {
        "min" : 0.0000000000000000,
        "max" : 1309937401.0000000000000000
    },
    "wrong_fragment" : {
        "min" : 0.0000000000000000,
        "max" : 3.0000000000000000
    },
    "urgent" : {
        "min" : 0.0000000000000000,
        "max" : 14.0000000000000000
    },
    "hot" : {
        "min" : 0.0000000000000000,
        "max" : 77.0000000000000000
    },
    "num_failed_logins" : {
        "min" : 0.0000000000000000,
        "max" : 5.0000000000000000
    },
    "num_compromised" : {
        "min" : 0.0000000000000000,
        "max" : 7479.0000000000000000
    },
    "root_shell" : {
        "min" : 0.0000000000000000,
        "max" : 1.0000000000000000
    },
    "su_attempted" : {
        "min" : 0.0000000000000000,
        "max" : 2.0000000000000000
    },
    "num_root" : {
        "min" : 0.0000000000000000,
        "max" : 7468.0000000000000000
    },
    "num_file_creations" : {
        "min" : 0.0000000000000000,
        "max" : 43.0000000000000000
    },
    "num_shells" : {
        "min" : 0.0000000000000000,
        "max" : 2.0000000000000000
    },
    "num_access_files" : {
        "min" : 0.0000000000000000,
        "max" : 9.0000000000000000
    },
    "num_outbound_cmds" : {
        "min" : 0.0000000000000000,
        "max" : 0.0000000000000000
    },
    "count" : {
        "min" : 0.0000000000000000,
        "max" : 511.0000000000000000
    },
    "srv_count" : {
        "min" : 0.0000000000000000,
        "max" : 511.0000000000000000
    },
    "serror_rate" : {
        "min" : 0.0000000000000000,
        "max" : 0.0000000000000000
    },
    "srv_serror_rate" : {
        "min" : 0.0000000000000000,
        "max" : 0.0000000000000000
    },
    "rerror_rate" : {
        "min" : 0.0000000000000000,
        "max" : 0.0000000000000000
    },
    "srv_rerror_rate" : {
        "min" : 0.0000000000000000,
        "max" : 0.0000000000000000
    },
    "same_srv_rate" : {
        "min" : 0.0000000000000000,
        "max" : 0.0000000000000000
    },
    "diff_srv_rate" : {
        "min" : 0.0000000000000000,
        "max" : 0.0000000000000000
    },
    "srv_diff_host_rate" : {
        "min" : 0.0000000000000000,
        "max" : 0.0000000000000000
    },
    "dst_host_count" : {
        "min" : 0.0000000000000000,
        "max" : 255.0000000000000000
    },
    "dst_host_srv_count" : {
        "min" : 0.0000000000000000,
        "max" : 255.0000000000000000
    },
    "dst_host_same_srv_rate" : {
        "min" : 0.0000000000000000,
        "max" : 0.0000000000000000
    },
    "dst_host_diff_srv_rate" : {
        "min" : 0.0000000000000000,
        "max" : 0.0000000000000000
    },
    "dst_host_same_src_port_rate" : {
        "min" : 0.0000000000000000,
        "max" : 0.0000000000000000
    },
    "dst_host_srv_diff_host_rate" : {
        "min" : 0.0000000000000000,
        "max" : 0.0000000000000000
    },
    "dst_host_serror_rate" : {
        "min" : 0.0000000000000000,
        "max" : 0.0000000000000000
    },
    "dst_host_srv_serror_rate" : {
        "min" : 0.0000000000000000,
        "max" : 0.0000000000000000
    },
    "dst_host_rerror_rate" : {
        "min" : 0.0000000000000000,
        "max" : 0.0000000000000000
    },
    "dst_host_srv_rerror_rate" : {
        "min" : 0.0000000000000000,
        "max" : 0.0000000000000000
    }
}
