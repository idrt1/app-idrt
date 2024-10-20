<script lang="ts">
    import Client from "../Client.svelte";
    import { page } from "$app/stores";
    import { client } from "$lib/wailsjs/go/models";



    let clients:client.Client | undefined;


</script>


<Client {clients} mode="create"/>