<script lang="ts">
	import { AppBar } from "@skeletonlabs/skeleton";
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Home" />
</svelte:head>

<div class="w-full h-full">
	<AppBar>
		<svelte:fragment slot="lead">Turnix</svelte:fragment>
		
		<svelte:fragment slot="trail">
			<a class="hover:underline" href="/z/pages/portal">Portal</a>
			<a class="hover:underline" href="/z/pages/auth/login">Login</a>
			<a class="hover:underline" href="/z/pages/auth/login">Signup</a>
		</svelte:fragment>


	</AppBar>
</div>
