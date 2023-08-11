<script lang="ts">
	import type { JWT } from '../../interfaces';
	import '../app.css';
	import Navbar from '../components/navbar.svelte';
	export let data: JWT;
</script>

<main>
	<Navbar data={data} />
	<slot />
</main>
