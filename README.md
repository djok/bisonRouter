# bisonRouter

apt install syslog-ng-core

/etc/syslog-ng/conf.d/zinc.conf
```
destination d_elasticsearch_http {
    elasticsearch-http(
        index("syslog-ng")
        type("")
        user("admin")
        password("admin")
        url("http://localhost:4080/api/_bulk")
        template("$(format-json --scope rfc5424 --scope dot-nv-pairs
        --rekey .* --shift 1 --scope nv-pairs
        --exclude DATE --key ISODATE @timestamp=${ISODATE})")
    );
};


log {
    source(src);
    destination(d_elasticsearch_http);
    flags(flow-control);
};
```