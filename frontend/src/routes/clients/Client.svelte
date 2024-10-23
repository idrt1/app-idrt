<script lang="ts">
    import { Button } from "$lib/components/ui/button";
    import { Card, CardContent, CardHeader, CardTitle } from "$lib/components/ui/card";
    import { Input } from "$lib/components/ui/input";
    import { Label } from "$lib/components/ui/label";
    import { Tabs, TabsContent, TabsList, TabsTrigger } from "$lib/components/ui/tabs";
    import { Textarea } from "$lib/components/ui/textarea";
    import Client from "./Personne.svelte";
    import Cabinet from "./Cabinet.svelte";
    import Syndicat from "./Syndicat.svelte";
    import * as Select from "$lib/components/ui/select";
    import { InsertClient } from "$lib/wailsjs/go/client/ClientMananger";
    import { client } from "$lib/wailsjs/go/models";
    import CalendarIcon from "lucide-svelte/icons/calendar";
    import {CalendarDate, DateFormatter, type DateValue, getLocalTimeZone} from "@internationalized/date";
    import { cn } from "$lib/utils.js";
    import { Calendar } from "$lib/components/ui/calendar/index.js";
    import * as Popover from "$lib/components/ui/popover/index.js";
    import { DeleteClientByID } from "$lib/wailsjs/go/client/ClientMananger";
    import { UpdateClient } from "$lib/wailsjs/go/client/ClientMananger";
    import * as AlertDialog from "$lib/components/ui/alert-dialog";
    import {afterUpdate} from "svelte";
    import {goto} from "$app/navigation";


    export let clients = new client.Client();
    export let mode : "update" | "create" = "create";

    let datePaiementValue: DateValue | undefined = undefined;

    const df = new DateFormatter("en-US", {
        dateStyle: "long"
    });

    const syndique = [
        { value: "N", label: "N" },
        { value: "D", label: "D" },
        {value: "S", label: "S"}
    ];

    function formatDate(date: Date): string {
        console.log(date.toISOString().split('T')[0]);
        return date.toISOString().split('T')[0];
    }

    function generate() {
        if (mode === "update") {
            UpdateClient(clients).then((result) => {
                goto("/");
            });
        }else{
            InsertClient(clients).then((result) => {
                console.log(result);
                goto("/");
            });
        }
    }
    async function deleteUser() {
        if (clients.id) {
            await DeleteClientByID(clients.id);
            goto("/");
        }
    }

    afterUpdate(() => {
        if (clients.datePaiement) {
            const datePaiement = new Date(clients.datePaiement);
            datePaiementValue = new CalendarDate(datePaiement.getFullYear(), datePaiement.getMonth() + 1, datePaiement.getDate());
        }
    });


    $: clients.datePaiement = datePaiementValue ? formatDate(datePaiementValue.toDate(getLocalTimeZone())) : "";
</script>


<div class="container mx-auto p-4 space-y-6">
    <div class="flex flex-row justify-between">
        <h1 class="text-3xl font-bold">Nouveau client</h1>
        <div class="flex space-x-4">
            <Button variant="outline" href="/">Accueil</Button>
            <AlertDialog.Root>
                <AlertDialog.Trigger asChild let:builder>
                    <Button builders={[builder]} variant="destructive">Supprimer</Button>
                </AlertDialog.Trigger>
                <AlertDialog.Content>
                    <AlertDialog.Header>
                        <AlertDialog.Title>Êtes-vous sûr de vouloir supprimer le client ?</AlertDialog.Title>
                        <AlertDialog.Description>
                            Cette action ne peut être annulée.
                            Elle supprimera définitivement le compte.
                        </AlertDialog.Description>
                    </AlertDialog.Header>
                    <AlertDialog.Footer>
                        <AlertDialog.Cancel>Retour</AlertDialog.Cancel>
                        <AlertDialog.Action on:click={deleteUser} class="bg-red-500 hover:bg-red-600 dark:bg-red-600 dark:hover:bg-red-700">Supprimer</AlertDialog.Action>
                    </AlertDialog.Footer>
                </AlertDialog.Content>
            </AlertDialog.Root>
        </div>
    </div>
    <Card>
        <CardHeader>
            <CardTitle>Informations principales</CardTitle>
        </CardHeader>
        <CardContent class="space-y-4">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div class="space-y-2 flex flex-col">
                    <div class="space-y-2">
                        <Label for="nom-type">Nom</Label>
                        <Input id="nom-type" placeholder="Entrée le Nom" bind:value={clients.nom}/>
                    </div>
                    <div class="space-y-2">
                        <Label for="prenom-type">Prénom</Label>
                        <Input id="prenom-type" placeholder="Entrée le prénom" bind:value={clients.prenom}/>
                    </div>
                    <div class="space-y-2">
                        <Label for="titre-type">Titre</Label>
                        <Input id="titre-type" placeholder="Entrée le titre" bind:value={clients.titre}/>
                    </div>
                </div>
                <div class="space-y-2 flex flex-col">
                    <div class="space-y-2">
                        <Label for="image">Image</Label>
                        <Input id="image" type="file" />
                    </div>
                    <div class="space-y-2">
                        <Label for="telephone-type">Téléphone</Label>
                        <Input id="telephone-type" placeholder="Entrée le numéro de téléphone" bind:value={clients.numeroPort}/>
                    </div>
                    <div class="space-y-2">
                        <Label for="email">Email</Label>
                        <Input type="email" id="email" placeholder="email" bind:value={clients.adresseElectronique}/>
                    </div>
                    <div class="space-y-2 flex flex-row">
                        <Popover.Root>
                            <div class="flex flex-col space-y-2 mr-3">
                                <Label for="date">Date de paiement</Label>
                                <Popover.Trigger asChild let:builder>
                                    <Button variant="outline" class={cn("w-[280px] justify-start text-left font-normal", !datePaiementValue && "text-muted-foreground")} builders={[builder]}>
                                        <CalendarIcon class="mr-2 h-4 w-4" />
                                        {datePaiementValue ? df.format(datePaiementValue.toDate(getLocalTimeZone())) : "Choisir une date"}
                                    </Button>
                                </Popover.Trigger>
                            </div>
                            <Popover.Content class="w-auto p-0">
                                <Calendar bind:value={datePaiementValue} initialFocus />
                            </Popover.Content>
                        </Popover.Root>
                        <div class="flex flex-col">
                            <Label for="syndique">Syndiqué</Label>
                            <Select.Root portal={null} selected={{
                                value:clients.syndique,
                                label:syndique.find(v=>v.value === clients.syndique)?.label
                            }}
                            onSelectedChange={(v) => v && (clients.syndique = v.value)}>
                                <Select.Trigger class="w-[180px]">
                                    <Select.Value placeholder="Syndiqué" />
                                </Select.Trigger>
                                <Select.Content>
                                    <Select.Group>
                                        <Select.Label>Syndiqué</Select.Label>
                                        {#each syndique as syndique}
                                            <Select.Item value={syndique.value} label={syndique.label}>
                                                {syndique.label}
                                            </Select.Item>
                                        {/each}
                                    </Select.Group>
                                </Select.Content>
                                <Select.Input name="syndique"/>
                            </Select.Root>
                        </div>
                    </div>
                </div>
            </div>
        </CardContent>
    </Card>


    <Tabs value="union" class="w-full">
        <TabsList class="grid w-full grid-cols-4">
            <TabsTrigger value="union" class="data-[state=active]:bg-[#EF4343]/30">Syndicat</TabsTrigger>
            <TabsTrigger value="firm" class="data-[state=active]:bg-[#facc14]/30">Cabinet</TabsTrigger>
            <TabsTrigger value="person" class="data-[state=active]:bg-[#1DAA51]/30">Personne</TabsTrigger>
            <TabsTrigger value="remark" class="data-[state=active]:bg-[#0da2e7]/30">Remarque</TabsTrigger>
        </TabsList>
        <TabsContent value="union" class="mt-4">
            <Syndicat bind:clients/>
        </TabsContent>
        <TabsContent value="firm" class="mt-4">
            <Cabinet bind:clients/>
        </TabsContent>
        <TabsContent value="person" class="mt-4">
            <Client bind:clients/>
        </TabsContent>
        <TabsContent value="remark" class="mt-4">
            <Card class="border-[#0da2e7]/50">
                <CardHeader>
                    <CardTitle>Remarque</CardTitle>
                </CardHeader>
                <CardContent>
                    <div class="space-y-2">
                        <Label for="remarks">Remarque supplémentaires</Label>
                        <Textarea id="remarks" placeholder="Entré des remarque supplémentaires" rows={4} bind:value={clients.remarques} />
                    </div>
                </CardContent>
            </Card>
        </TabsContent>
    </Tabs>

    <div class="flex justify-end space-x-4">
        <Button variant="outline" href="/">Annuler</Button>
        <Button on:click={generate}>Enregistrer le client</Button>
    </div>
</div>