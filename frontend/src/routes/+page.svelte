<script lang="ts">
    import { onMount } from "svelte";
    import { Button } from "$lib/components/ui/button";
    import { PlusCircle } from "lucide-svelte"
    import { GetAllClient } from "$lib/wailsjs/go/client/ClientMananger";
    import { client } from "$lib/wailsjs/go/models";
    import DataTable from "./DataTable/DataTable.svelte";

    let clients: client.Client[] = [];

    onMount(async() => {
        try {
            clients = await GetAllClient()
        } catch(err) {
            console.error(err)
        }
    })
</script>


<div class="min-h-screen bg-background flex flex-col items-center justify-center p-4">
    <main class="w-full">
        <h1 class="text-3xl font-bold text-center text-foreground">SCD-17</h1>
        <div class="space-y-6">
            <div class="space-y-4 mx-auto text-center pt-4 pl-8" style="max-width: 1000px;">
                <Button class="w-full" size="lg" href="/clients">
                    <PlusCircle class="mr-2 h-5 w-5" />
                    Ajouter un client
                </Button>
            </div>
            {#if clients.length !== 0}
                <div class="w-full overflow-x-auto mx-auto" style="max-width: 1000px;">
                    <DataTable clients={clients}/>
                </div>
            {:else}
                <p class="text-center text-muted-foreground">Aucun client trouvé</p>
            {/if}
        </div>
    </main>
</div>



