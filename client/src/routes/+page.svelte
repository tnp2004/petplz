<script lang="ts">
	import { onMount } from 'svelte';
	import { authenticated } from '../stores/auth';

	let message = '';

	onMount(async () => {
			const response = await fetch('http://localhost:3000/api/accounts', {
				headers: { 'Content-Type': 'application/json' },
				credentials: 'include'
			});

            const resOkay = response.ok
            if (resOkay) {
                const content = await response.json();
                authenticated.set(true);
                message = `Hi, ${content.username}`;
                return
            }
            message = "you are not logged in"
            
            authenticated.set(resOkay);

	});
</script>

<h1>petplz Homepage</h1>
{message}
