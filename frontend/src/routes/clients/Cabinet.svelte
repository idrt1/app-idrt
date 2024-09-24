<script>
    import {Card, CardContent, CardHeader, CardTitle} from "$lib/components/ui/card";
    import {Input} from "$lib/components/ui/input";
    import {Label} from "$lib/components/ui/label";
</script>

<Card class="border-[#facc14]/50">
    <CardHeader>
        <CardTitle>Cabinet information</CardTitle>
    </CardHeader>
    <CardContent class="space-y-4">
        <div class="space-y-2">
            <Label for="cabinet-lieux">Lieux d'exercice</Label>
            <Input id="cabinet-lieux" placeholder="Entré un lieux" />
        </div>
        <div class="space-y-2">
            <Label for="cabinet-type">Type d'exercice</Label>
            <Input id="cabinet-type" placeholder="Entré un type" />
        </div>
        <div class="space-y-2">
                <Label for="cabinet-application">Application</Label>
            <Input id="cabinet-application" placeholder="Entré une application" />
        </div>
    </CardContent>
</Card>