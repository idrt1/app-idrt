<script lang="ts">
    import {Card, CardContent, CardHeader, CardTitle} from "$lib/components/ui/card";
    import {Input} from "$lib/components/ui/input";
    import {Label} from "$lib/components/ui/label";
    import CalendarIcon from "lucide-svelte/icons/calendar";
    import { DateFormatter, type DateValue, getLocalTimeZone } from "@internationalized/date";
    import { cn } from "$lib/utils.js";
    import { Calendar } from "$lib/components/ui/calendar/index.js";
    import * as Popover from "$lib/components/ui/popover/index.js";
    import { Button } from "$lib/components/ui/button";
    import {Checkbox} from "$lib/components/ui/checkbox";
    import * as Select from "$lib/components/ui/select";
    import {client} from "$lib/wailsjs/go/models";

    const df = new DateFormatter("en-US", { dateStyle: "long" });

    let dateCreationValue: DateValue | undefined = undefined;
    let dateInstallationValue: DateValue | undefined = undefined;

    let checkedYes: boolean = false;
    let checkedNo: boolean = false;

    export let clients : client.Client;

    const installation = [
        { value: "Libéral", label: "Libéral" },
        { value: "Collaborateur", label: "Collaborateur" },
        { value: "Salarié", label: "Salarié" },
        { value: "Retraité", label: "Retraité" },
        { value: "Etudiant", label: "Etudiant" },
        { value: "Mutuelle", label: "Mutuelle" },
        { value: "Contrôlleur", label: "Contrôlleur" },
    ];

    const categorie = [
        { value : "A", label: "A" },
        { value : "A1/2", label: "A1/2" },
        { value : "CA", label: "C" },
        { value : "CS", label: "D" },
        { value : "CS1/2", label: "E" },
        { value : "etd", label: "C" },
        { value : "P35", label: "D" },
        { value : "P351/2", label: "E" },
    ];

    function formatDate(date: Date): string {
        const options: Intl.DateTimeFormatOptions = { day: '2-digit', month: '2-digit', year: 'numeric' };
        return date.toLocaleDateString('fr-FR', options);
    }

    $:clients.dateCreation = dateCreationValue ? formatDate(dateCreationValue.toDate(getLocalTimeZone())) : "";
    $:clients.dateInstallation = dateInstallationValue ? formatDate(dateInstallationValue.toDate(getLocalTimeZone())) : "";

    $: clients.conjointSynd = checkedYes ? "O" : checkedNo ? "N" : "";
</script>

<Card class="border-[#EF4343]/50">
    <CardHeader>
        <CardTitle>Information du syndicat</CardTitle>
    </CardHeader>
    <CardContent class="space-y-4">
        <div class="space-y-2">
            <Label for="start-date">Date de création</Label>
            <Popover.Root>
                <Popover.Trigger asChild let:builder>
                    <Button variant="outline" class={cn("h-[42px] w-full justify-start text-left font-normal", !dateCreationValue && "text-muted-foreground")} builders={[builder]}>
                        <CalendarIcon class="mr-2 h-4 w-4" />
                        {dateCreationValue ? df.format(dateCreationValue.toDate(getLocalTimeZone())) : "Choisir une date de création"}
                    </Button>
                </Popover.Trigger>
                <Popover.Content class="w-auto p-0">
                    <Calendar bind:value={dateCreationValue} initialFocus />
                </Popover.Content>
            </Popover.Root>
        </div>
        <div class="space-y-2">
            <Label for="end-date">Date d'installation</Label>
            <Popover.Root>
                <Popover.Trigger asChild let:builder>
                    <Button variant="outline" class={cn("h-[42px] w-full justify-start text-left font-normal", !dateInstallationValue && "text-muted-foreground")} builders={[builder]}>
                        <CalendarIcon class="mr-2 h-4 w-4" />
                        {dateInstallationValue ? df.format(dateInstallationValue.toDate(getLocalTimeZone())) : "Choisir une date d'installation"}
                    </Button>
                </Popover.Trigger>
                <Popover.Content class="w-auto p-0">
                    <Calendar bind:value={dateInstallationValue} initialFocus />
                </Popover.Content>
            </Popover.Root>
        </div>

        <div class="space-y-2">
            <Label for="syndicat-coti">Premier an coti</Label>
            <Input id="syndicat-coti" placeholder="Entré un premier an" bind:value={clients.premierAnCoti} />
        </div>

        <div id="conjoint-syndic" class="flex flex-col space-x-2 space-y-2">
            <Label for="syndicat-coti">Conjoint syndic</Label>
            <div class="flex items-center space-x-2">
                <Checkbox id="terms-yes" bind:checked={checkedYes} aria-labelledby="terms-label-yes" />
                <Label id="terms-label-yes" for="terms-yes" class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70">
                    O
                </Label>
                <Checkbox id="terms-no" bind:checked={checkedNo} aria-labelledby="terms-label-no" />
                <Label id="terms-label-no" for="terms-no" class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70">
                    N
                </Label>
            </div>
        </div>
        <div class="space-y-2">
            <Label for="installation">Installation</Label>
            <Select.Root portal={null} selected={{
                                value:clients.typeInstallation,
                                label:installation.find(v=>v.value === clients.typeInstallation)?.label
                            }}
                         onSelectedChange={(v) => v && (clients.typeInstallation = v.value)}>
                <Select.Trigger class="w-full">
                    <Select.Value placeholder="Type d'installation" />
                </Select.Trigger>
                <Select.Content>
                    <Select.Group>
                        <Select.Label>Installation</Select.Label>
                        {#each installation as installation}
                            <Select.Item value={installation.value} label={installation.label}>
                                {installation.label}
                            </Select.Item>
                        {/each}
                    </Select.Group>
                </Select.Content>
                <Select.Input name="installation" />
            </Select.Root>
        </div>
        <div class="space-y-2">
            <Label for="installation">Categorie</Label>
            <Select.Root portal={null} selected={{
                                value:clients.categorie,
                                label:categorie.find(v=>v.value === clients.categorie)?.label
                            }}
                         onSelectedChange={(v) => v && (clients.categorie = v.value)}>
                <Select.Trigger class="w-full">
                    <Select.Value placeholder="Categorie" />
                </Select.Trigger>
                <Select.Content>
                    <Select.Group>
                        <Select.Label>Categorie</Select.Label>
                        {#each categorie as categorie}
                            <Select.Item value={categorie.value} label={categorie.label}>
                                {categorie.label}
                            </Select.Item>
                        {/each}
                    </Select.Group>
                </Select.Content>
                <Select.Input name="categorie" />
            </Select.Root>
        </div>
        <div class="space-y-2">
            <Label for="syndicat-number">N° Union</Label>
            <Input id="syndicat-number" placeholder="Entrer un numéros d'union" bind:value={clients.idSyndique} />
        </div>
        <div class="space-y-2">
            <Label for="syndicat-cotisation">Cotisation</Label>
            <Input id="syndicat-cotisation" placeholder="Entrer une cotisation" bind:value={clients.cotisation} />
        </div>
        <div class="space-y-2">
            <Label for="syndicat-cout">Coût</Label>
            <Input id="syndicat-cout" placeholder="Entrer un coût" bind:value={clients.cout} />
        </div>
        <div class="space-y-2">
            <Label for="syndicat-cotisationSpé">Cotisation spéciale</Label>
            <Input id="syndicat-cotisationSpé" placeholder="Entrer une cotisation spé" bind:value={clients.cotiSpeciale} />
        </div>
    </CardContent>
</Card>
