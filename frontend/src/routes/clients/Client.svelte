<script>
    import {Card, CardContent, CardHeader, CardTitle} from "$lib/components/ui/card";
    import {Label} from "$lib/components/ui/label";
    import {Input} from "$lib/components/ui/input";
    import * as Select from "$lib/components/ui/select";
    const sexe = [
        { value: "M", label: "M" },
        { value: "F", label: "F" }
    ];
</script>

<Card class="border-[#1DAA51]/50">
    <CardHeader>
        <CardTitle>Information du client</CardTitle>
    </CardHeader>
    <CardContent class="space-y-4">
        <div class="space-y-2">
            <Label for="person-batiment">Bâtiment</Label>
            <Input id="person-batiment" placeholder="Enter building" />
        </div>
        <div class="space-y-2">
            <Label for="person-address">Adresse</Label>
            <Input id="person-address" placeholder="Enter address" />
        </div>
        <div class="space-y-2">
            <Label for="person-postal-code">Code Postal</Label>
            <Input id="person-postal-code" placeholder="Enter postal code" />
        </div>
        <div class="space-y-2">
            <Label for="person-city">Ville</Label>
            <Input id="person-city" placeholder="Enter city" />
        </div>
        <div class="space-y-2">
            <Label for="person-country">Pays</Label>
            <Input id="person-country" placeholder="Enter country" />
        </div>
        <div class="space-y-2">
            <Label for="person-dob">Date de naissance</Label>
            <Input id="person-dob" type="date" />
        </div>
        <div class="flex flex-col space-y-2">
            <Label for="sexe">Sexe</Label>
            <Select.Root portal={null}>
                <Select.Trigger class="w-[180px]">
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
    </CardContent>
</Card>
