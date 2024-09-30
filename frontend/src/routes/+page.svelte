<script lang="ts">
    import { Button } from "$lib/components/ui/button";
    import { Input } from "$lib/components/ui/input"
    import { PlusCircle, Search } from "lucide-svelte"
</script>


<div class="min-h-screen bg-background flex flex-col items-center justify-center p-4">
    <main class="max-w-md w-full space-y-6">
        <h1 class="text-3xl font-bold text-center text-foreground">SCD-17</h1>
        <div class="space-y-4">
            <Button class="w-full" size="lg" href="/clients">
                <PlusCircle class="mr-2 h-5 w-5" />
                Ajouter un client
            </Button>
            <div class="relative">
                <Input
                        type="search"
                        placeholder="Rechercher un client"
                        class="w-full pl-10 pr-4 py-2"
                />
                <Search class="absolute left-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-muted-foreground" />
            </div>
        </div>
    </main>
</div>