<script lang="ts">
    import { Card, CardContent, CardHeader, CardTitle } from "$lib/components/ui/card";
    import { Label } from "$lib/components/ui/label";
    import { Input } from "$lib/components/ui/input";
    import * as Select from "$lib/components/ui/select";
    const sexe = [
        { value: "M", label: "M" },
        { value: "F", label: "F" }
    ];
    import CalendarIcon from "lucide-svelte/icons/calendar";
    import {
        DateFormatter,
        type DateValue,
        getLocalTimeZone
    } from "@internationalized/date";
    import { cn } from "$lib/utils.js";
    import { Calendar } from "$lib/components/ui/calendar/index.js";
    import * as Popover from "$lib/components/ui/popover/index.js";
    import { Button } from "$lib/components/ui/button";
    import {Checkbox} from "$lib/components/ui/checkbox";
    import {client} from "$lib/wailsjs/go/models";

    const df = new DateFormatter("en-US", { dateStyle: "long" });
    let value: DateValue | undefined = undefined;

    function allowOnlyNumbers(event: InputEvent) {
        const input = event.target as HTMLInputElement;
        input.value = input.value.replace(/\D/g, '');
    }

    export let clients : client.Client;

</script>

<Card class="border-[#1DAA51]/50">
    <CardHeader>
        <CardTitle>Information du client</CardTitle>
    </CardHeader>
    <CardContent class="space-y-4">
        <div class="flex flex-row space-x-4">
            <div class="flex flex-col w-1/2 space-y-4">
                <div class="space-y-2">
                    <Label for="person-batiment">Bâtiment</Label>
                    <Input id="person-batiment" placeholder="Entré le bâtiment" bind:value={clients.adresseProf1} />
                </div>
                <div class="space-y-2">
                    <Label for="person-address">Adresse</Label>
                    <Input id="person-address" placeholder="Entré l'adresse" bind:value={clients.adresseProf2} />
                </div>
                <div class="space-y-2">
                    <Label for="person-postal-code">Code Postal</Label>
                    <Input id="person-postal-code" placeholder="Entré le code postal" bind:value={clients.codePostalProf} />
                </div>
                <div class="space-y-2">
                    <Label for="person-city">Ville</Label>
                    <Input id="person-city" placeholder="Entré la ville" bind:value={clients.villeProf} />
                </div>
                <div class="space-y-2">
                    <Label for="person-tel-pro">Tél professionnel</Label>
                    <Input id="person-tel-pro" type="tel" placeholder="Entré le numéro professionnel" on:input={allowOnlyNumbers} bind:value={clients.numeroTelProf} />
                </div>
                <div class="space-y-2">
                    <Label for="person-tel-portable">Tél portable</Label>
                    <Input id="person-tel-portable" type="tel" placeholder="Entré le numéro portable" on:input={allowOnlyNumbers} bind:value={clients.numeroTelDomicile} />
                </div>
            </div>

            <div class="flex flex-col w-1/2 space-y-4">
                <div class="space-y-2">
                    <Label for="person-country">Pays</Label>
                    <Input id="person-country" placeholder="Entré le pays" bind:value={clients.paysProf} />
                </div>
                <div class="space-y-2">
                    <Popover.Root>
                        <div class="flex flex-col space-y-2">
                            <Label for="date">Date de paiement</Label>
                            <Popover.Trigger asChild let:builder>
                                <Button variant="outline" class={cn("h-[42px] w-full justify-start text-left font-normal", !value && "text-muted-foreground")} builders={[builder]}>
                                    <CalendarIcon class="mr-2 h-4 w-4" />
                                    {value ? df.format(value.toDate(getLocalTimeZone())) : "Choisir une date"}
                                </Button>
                            </Popover.Trigger>
                        </div>
                        <Popover.Content class="w-auto p-0">
                            <Calendar bind:value initialFocus />
                        </Popover.Content>
                    </Popover.Root>
                </div>
                <div class="space-y-2">
                    <Label for="person-diplome">Diplôme faculté</Label>
                    <Input id="person-diplome" placeholder="Entré le diplôme" />
                </div>
                <div class="space-y-2">
                    <Popover.Root>
                        <div class="flex flex-col space-y-2">
                            <Label for="date">Date du diplôme</Label>
                            <Popover.Trigger asChild let:builder>
                                <Button variant="outline" class={cn("h-[42px] w-full justify-start text-left font-normal", !value && "text-muted-foreground")} builders={[builder]}>
                                    <CalendarIcon class="mr-2 h-4 w-4" />
                                    {value ? df.format(value.toDate(getLocalTimeZone())) : "Choisir une date"}
                                </Button>
                            </Popover.Trigger>
                        </div>
                        <Popover.Content class="w-auto p-0">
                            <Calendar bind:value initialFocus />
                        </Popover.Content>
                    </Popover.Root>
                </div>
                <div class="space-y-2">
                    <Label for="sexe">Sexe</Label>
                    <Select.Root portal={null}>
                        <Select.Trigger class="w-full">
                            <Select.Value placeholder="Sexe" />
                        </Select.Trigger>
                        <Select.Content>
                            <Select.Group>
                                <Select.Label>Sexe</Select.Label>
                                {#each sexe as sexe}
                                    <Select.Item value={sexe.value} label={sexe.label}>
                                        {sexe.label}
                                    </Select.Item>
                                {/each}
                            </Select.Group>
                        </Select.Content>
                        <Select.Input name="sexe" />
                    </Select.Root>
                </div>
                <div class="space-y-2">
                    <Label for="person-tel-domicile">Tél domicile</Label>
                    <Input id="person-tel-domicile" type="tel" placeholder="Entré le numéro domicile" on:input={allowOnlyNumbers} />
                </div>
                <div class="space-y-2">
                    <Label for="person-tel-fax">Tél fax</Label>
                    <Input id="person-tel-fax" type="tel" placeholder="Entré le numéro fax" on:input={allowOnlyNumbers} />
                </div>
            </div>
        </div>
    </CardContent>
</Card>
