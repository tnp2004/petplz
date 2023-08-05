<script lang="ts">
	import { onMount } from 'svelte';
	import { authenticated } from '../stores/auth';
	import { PUBLIC_SERVER_URL } from '$env/static/public';

	let message = '';

	onMount(async () => {
			const response = await fetch(`${PUBLIC_SERVER_URL}/api/accounts`, {
				headers: { 'Content-Type': 'application/json' },
				credentials: 'include'
			});

            const resOkay = response.ok
            authenticated.set(resOkay);
            if (resOkay) {
                const data = await response.json();
                message = `Hi, ${data.username}`;
                return
            }
            message = "you are not logged in"
            

	});
</script>

<h1>petplz Homepage</h1>
{message}
