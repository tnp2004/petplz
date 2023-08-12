<script lang="ts">
	import type { Account } from '../../../../interfaces';
	import type { PageData } from './$types';
    export let data: PageData;

    var gender = {
        "male": "ชาย",
        "female": "หญิง",
        "other": "อื่นๆ"
    }

	let accountData: Account = data.account;
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
