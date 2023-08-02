<script lang="ts">
	import { onMount } from 'svelte';
    import { authenticated } from "../../../stores/auth.ts"

    interface Account {
        accountId: string,
        username: string,
        email: string,
        gender: "male" | "female" | "other",
        age: number,
        money: number,
        image_uri: string,
    }

    let accountData: Account;

    onMount(async () => {
			const response = await fetch('http://localhost:3000/api/accounts', {
				headers: { 'Content-Type': 'application/json' },
				credentials: 'include'
			});

            const resOkay = response.ok
            authenticated.set(resOkay);
            if (response.ok) {
                const data = await response.json();
                accountData = data
            }
	});
</script>
{#if accountData}
    {JSON.stringify(accountData)}
{/if}