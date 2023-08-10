<script lang="ts">
	import { onMount } from 'svelte';
	import { authenticated } from '../../../stores/auth';
	import { PUBLIC_SERVER_URL } from '$env/static/public';
	import type { Account } from '../../../../interfaces';

    var gender = {
        "male": "ชาย",
        "female": "หญิง",
        "other": "อื่นๆ"
    }

	let accountData: Account;

	onMount(async () => {
		const response = await fetch(`${PUBLIC_SERVER_URL}/api/accounts`, {
			headers: { 'Content-Type': 'application/json' },
			credentials: 'include'
		});

		const resOkay = response.ok;
		authenticated.set(resOkay);
		if (response.ok) {
			const data = await response.json();
			accountData = data;
		}
	});
</script>

{#if accountData}
	<div class="w-fit px-5 mx-auto">
        <div class="text-xl">
            <div class="text-center my-3">
                <label class="font-bold text-4xl" for="profile">Profile</label>
            </div>
            <div class="mx-auto my-3 w-64 h-64 rounded-full overflow-hidden">
                <!-- svelte-ignore a11y-img-redundant-alt -->
                <img src={accountData.image_uri} alt="profile-image" />
            </div>
            <div class="flex justify-between">
                <label class="font-bold w-1/2" for="username">ชื่อผู้ใช้</label>
                <div class="w-1/2">
                    {accountData.username}
                </div>
            </div>
            <div class="flex justify-between">
                <label class="font-bold w-1/2" for="email">อีเมล</label>
                <div class="w-1/2">
                    {accountData.email}
                </div>
            </div>
            <div class="flex justify-between">
                <label class="font-bold w-1/2" for="gender">เพศ</label>
                <div class="w-1/2">
                    {gender[accountData.gender]}
                </div>
            </div>
            <div class="flex justify-between">
                <label class="font-bold w-1/2" for="age">อายุ</label>
                <div class="w-1/2">
                    {accountData.age}
                </div>
            </div>
           
        </div>
    </div>
{/if}
