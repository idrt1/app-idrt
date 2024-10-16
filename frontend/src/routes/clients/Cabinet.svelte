<script lang="ts">
    import { Card, CardContent, CardHeader, CardTitle } from "$lib/components/ui/card";
    import { Input } from "$lib/components/ui/input";
    import { Label } from "$lib/components/ui/label";
    import { Checkbox } from "$lib/components/ui/checkbox";
    import { client } from "$lib/wailsjs/go/models";

    let checkedYes = false;
    let checkedNo = false;
    let roles = {
        femmeDeMenage: false,
        receptioniste: false,
        assistante: false,
        collaborateur: false,
        laboratoire: false,
        aide: false
    };

    $: clients.d_ou_N = checkedYes ? "O" : checkedNo ? "N" : "";

    function convertRolesToNumbers() {
        clients.persAideAssistante = roles.aide ? "1" : "0";
        clients.persCollaborateur = roles.collaborateur ? "1" : "0";
        clients.persAssistante = roles.assistante ? "1" : "0";
        clients.persReceptionniste = roles.receptioniste ? "1" : "0";
        clients.persFemmeDeMenage = roles.femmeDeMenage ? "1" : "0";
        clients.persLaboratoire = roles.laboratoire ? "1" : "0";
    }

    $: convertRolesToNumbers();
    export let clients : client.Client;

</script>

<Card>
    <CardHeader>
        <CardTitle>Information du cabinet</CardTitle>
    </CardHeader>
    <CardContent>
        <div>
            <Label>Lieux d'exercice</Label>
            <Input bind:value={clients.lieuExercice} />
        </div>
        <div>
            <Label>Type d'exercice</Label>
            <Input bind:value={clients.typeExercice} />
        </div>
        <div>
            <Label>Association</Label>
            <Input bind:value={clients.association} />
        </div>

        <div id="conjoint-syndic" class="flex flex-col space-x-2 space-y-2">
            <Label for="Personels">Personels</Label>
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
