<script>
    import {Card, CardContent, CardHeader, CardTitle} from "$lib/components/ui/card";
    import {Input} from "$lib/components/ui/input";
    import {Label} from "$lib/components/ui/label";
    import {Checkbox} from "$lib/components/ui/checkbox";

    let checked = false;
    let roles = {
        femmeDeMenage: false,
        receptioniste: false,
        assistante: false,
        collaborateur: false,
        laboratoire: false,
        aide: false
    };
</script>

<Card class="border-[#facc14]/50">
    <CardHeader>
        <CardTitle>Information du cabinet</CardTitle>
    </CardHeader>
    <CardContent class="space-y-4">
        <div class="space-y-2">
            <Label for="cabinet-lieux">Lieux d'exercice</Label>
            <Input id="cabinet-lieux" placeholder="Entré un lieu" />
        </div>
        <div class="space-y-2">
            <Label for="cabinet-type">Type d'exercice</Label>
            <Input id="cabinet-type" placeholder="Entré un type" />
        </div>
        <div class="space-y-2">
            <Label for="cabinet-association">Association</Label>
            <Input id="cabinet-association" placeholder="Entré une association" />
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
            <Label>Rôles disponibles</Label>
            <div class="flex flex-row justify-between">
                <div class="flex flex-col space-y-2">
                    <div class="flex items-center space-x-2">
                        <Checkbox id="femmeDeMenage" bind:checked={roles.femmeDeMenage} />
                        <Label for="femmeDeMenage" class="text-sm font-medium">Femme de ménage</Label>
                    </div>
                    <div class="flex items-center space-x-2">
                        <Checkbox id="receptioniste" bind:checked={roles.receptioniste} />
                        <Label for="receptioniste" class="text-sm font-medium">Réceptionniste</Label>
                    </div>
                </div>
                <div class="flex flex-col space-y-2">
                    <div class="flex items-center space-x-2">
                        <Checkbox id="assistante" bind:checked={roles.assistante} />
                        <Label for="assistante" class="text-sm font-medium">Assistante</Label>
                    </div>
                    <div class="flex items-center space-x-2">
                        <Checkbox id="collaborateur" bind:checked={roles.collaborateur} />
                        <Label for="collaborateur" class="text-sm font-medium">Collaborateur</Label>
                    </div>
                </div>
                <div class="flex flex-col space-y-2">
                    <div class="flex items-center space-x-2">
                        <Checkbox id="laboratoire" bind:checked={roles.laboratoire} />
                        <Label for="laboratoire" class="text-sm font-medium">Laboratoire</Label>
                    </div>
                    <div class="flex items-center space-x-2">
                        <Checkbox id="aide" bind:checked={roles.aide} />
                        <Label for="aide" class="text-sm font-medium">Aide assistante</Label>
                    </div>
                </div>
            </div>
        </div>
    </CardContent>
</Card>
