
<script>
	import  {AppShell} from "@skeletonlabs/skeleton"
</script>

<AppShell>
	<svelte:fragment slot="header">Header</svelte:fragment>
	<svelte:fragment slot="sidebarLeft">Sidebar Left</svelte:fragment>
	<slot />

	<svelte:fragment slot="pageFooter">Page Footer</svelte:fragment>
</AppShell>