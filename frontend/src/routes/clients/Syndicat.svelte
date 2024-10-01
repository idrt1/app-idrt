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
    let startDate: DateValue | undefined = undefined;
    let endDate: DateValue | undefined = undefined;

    let checked: boolean = false;

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
                    <Button variant="outline" class={cn("h-[42px] w-full justify-start text-left font-normal", !startDate && "text-muted-foreground")} builders={[builder]}>
                        <CalendarIcon class="mr-2 h-4 w-4" />
                        {startDate ? df.format(startDate.toDate(getLocalTimeZone())) : "Choisir une date de création"}
                    </Button>
                </Popover.Trigger>
                <Popover.Content class="w-auto p-0">
                    <Calendar bind:value={startDate} initialFocus />
                </Popover.Content>
            </Popover.Root>
        </div>
        <div class="space-y-2">
            <Label for="end-date">Date d'installation</Label>
            <Popover.Root>
                <Popover.Trigger asChild let:builder>
                    <Button variant="outline" class={cn("h-[42px] w-full justify-start text-left font-normal", !endDate && "text-muted-foreground")} builders={[builder]}>
                        <CalendarIcon class="mr-2 h-4 w-4" />
                        {endDate ? df.format(endDate.toDate(getLocalTimeZone())) : "Choisir une date d'installation"}
                    </Button>
                </Popover.Trigger>
                <Popover.Content class="w-auto p-0">
                    <Calendar bind:value={endDate} initialFocus />
                </Popover.Content>
            </Popover.Root>
        </div>

        <div class="space-y-2">
            <Label for="syndicat-coti">Premier an coti</Label>
            <Input id="syndicat-coti" placeholder="Entré un premier an" bind:value={clients.premierAnCoti} />
        </div>
        <div class="flex items-center space-x-2">
            <Checkbox id="terms-yes" bind:checked aria-labelledby="terms-label-yes" />
            <Label id="terms-label-yes" for="terms-yes" class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70">
                O
            </Label>
            <Checkbox id="terms-no" bind:checked aria-labelledby="terms-label-no" />
            <Label id="terms-label-no" for="terms-no" class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70">
                N
            </Label>
        </div>
        <div class="space-y-2">
            <Label for="installation">Installation</Label>
            <Select.Root portal={null}>
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
            <Select.Root portal={null}>
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
            <Label for="syndicat-number">Union Number</Label>
            <Input id="syndicat-number" placeholder="Enter union number" />
        </div>
    </CardContent>
</Card>
