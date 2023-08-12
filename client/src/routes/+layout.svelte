<script lang="ts">
	import '../app.css';
	import Navbar from '../components/navbar.svelte';
	import type { LayoutData } from './$types';
	
	export let data: LayoutData;
</script>

<main>
	<Navbar data={data} />
	<slot />
</main>
