<script lang="ts">
    // {"mtype":"response","data":{"os_info":{"Runtime":"linux","Arch":"amd64","Name":"Manjaro Linux","Version":"rolling","Windows":{"Build":""},"Linux":{"Distro":"arch","Kernel":"6.12.4-1-MANJARO","PkgManager":"pacman"},"BSD":{"Kernel":"","PkgManager":""},"MacOs":{"VersionName":""}},"hostname":"dark-horse","memory":{"Total":32671809536,"Used":14445658112,"Buffers":2190561280,"Cached":8103108608,"Free":5057646592,"Available":18226151424,"Active":16497135616,"Inactive":6263427072,"SwapTotal":4294963200,"SwapUsed":0,"SwapCached":0,"SwapFree":4294963200,"Mapped":2938028032,"Shmem":318820352,"Slab":4108754944,"PageTables":124764160,"Committed":39054004224,"VmallocUsed":179089408,"MemAvailableEnabled":true}}}

    interface Props {
        data: string;
    }

    let { data }: Props = $props();
</script>

<!-- base64 decode and show image -->

<img src="data:image/png;base64,{data}" alt="Screenshot" />
