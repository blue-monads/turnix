<script lang="ts">
	import { AppBar } from "@skeletonlabs/skeleton";
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Home" />
</svelte:head>

<AppBar>
	
	<svelte:fragment slot="headline">Turnis</svelte:fragment>
	<svelte:fragment slot="trail">
		<a href="/z/pages/portal">Portal</a>
		<a href="/z/pages/auth/login">Login</a>
		<a href="/z/pages/auth/login">Signup</a>
	</svelte:fragment>


</AppBar>
