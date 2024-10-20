<script lang="ts">
    import Client from "./Client.svelte";
    import { page } from "$app/stores";
    import { client } from "$lib/wailsjs/go/models";
    import {onMount} from "svelte";
    import {GetClientByID} from "$lib/wailsjs/go/client/ClientMananger";


    let clients:client.Client | undefined;

    onMount(async () => {
        const id = $page.url.searchParams.get('id') || '';
        if (id) {
            clients = await GetClientByID(parseInt(id));
        }
    });

</script>


<Client {clients} mode="update"/>