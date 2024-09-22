<script>
    import { Button } from "$lib/components/ui/button";
    import { Card, CardContent, CardHeader, CardTitle } from "$lib/components/ui/card"
    import { Input } from "$lib/components/ui/input"
    import { Label } from "$lib/components/ui/label"
    import { Tabs, TabsContent, TabsList, TabsTrigger } from "$lib/components/ui/tabs"
    import { Textarea } from "$lib/components/ui/textarea"
</script>

<div class="container mx-auto p-4 space-y-6">
    <h1 class="text-3xl font-bold">Nouveau client</h1>

    <Card>
        <CardHeader>
            <CardTitle>Informations principales</CardTitle>
        </CardHeader>
        <CardContent class="space-y-4">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div class="space-y-2 flex flex-col">
                    <div class="space-y-2">
                        <Label for="nom-type">Nom</Label>
                        <Input id="nom-type" placeholder="Entrée le Nom" />
                    </div>
                    <div class="space-y-2">
                        <Label for="prenom-type">Prénom</Label>
                        <Input id="prenom-type" placeholder="Entrée le prénom" />
                    </div>
                    <div class="space-y-2">
                        <Label for="titre-type">Titre</Label>
                        <Input id="titre-type" placeholder="Entrée le titre" />
                    </div>
                </div>

                <div class="grid w-full max-w-sm items-center gap-1.5">
                    <Label for="image">Image</Label>
                    <Input id="image" type="file" />
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
            <Card class="border-[#EF4343]/50">
                <CardHeader>
                    <CardTitle>Union Information</CardTitle>
                </CardHeader>
                <CardContent class="space-y-4">
                    <div class="space-y-2">
                        <Label for="union-name">Union Name</Label>
                        <Input id="union-name" placeholder="Enter union name" />
                    </div>
                    <div class="space-y-2">
                        <Label for="union-number">Union Number</Label>
                        <Input id="union-number" placeholder="Enter union number" />
                    </div>
                </CardContent>
            </Card>
        </TabsContent>
        <TabsContent value="firm" class="mt-4">
            <Card class="border-[#facc14]/50">
                <CardHeader>
                    <CardTitle>Firm Information</CardTitle>
                </CardHeader>
                <CardContent class="space-y-4">
                    <div class="space-y-2">
                        <Label for="firm-name">Firm Name</Label>
                        <Input id="firm-name" placeholder="Enter firm name" />
                    </div>
                    <div class="space-y-2">
                        <Label for="firm-address">Firm Address</Label>
                        <Input id="firm-address" placeholder="Enter firm address" />
                    </div>
                </CardContent>
            </Card>
        </TabsContent>
        <TabsContent value="person" class="mt-4">
            <Card class="border-[#1DAA51]/50">
                <CardHeader>
                    <CardTitle>Person Information</CardTitle>
                </CardHeader>
                <CardContent class="space-y-4">
                    <div class="space-y-2">
                        <Label for="person-name">Name</Label>
                        <Input id="person-name" placeholder="Enter person's name" />
                    </div>
                    <div class="space-y-2">
                        <Label for="person-role">Role</Label>
                        <Input id="person-role" placeholder="Enter person's role" />
                    </div>
                </CardContent>
            </Card>
        </TabsContent>
        <TabsContent value="remark" class="mt-4">
            <Card class="border-[#0da2e7]/50">
                <CardHeader>
                    <CardTitle>Remarque</CardTitle>
                </CardHeader>
                <CardContent>
                    <div class="space-y-2">
                        <Label for="remarks">Remarque supplémentaires</Label>
                        <Textarea id="remarks" placeholder="Entré des remarque supplémentaires" rows={4} />
                    </div>
                </CardContent>
            </Card>
        </TabsContent>
    </Tabs>

    <div class="flex justify-end space-x-4">
        <Button variant="outline">Cancel</Button>
        <Button>Save Client</Button>
    </div>
</div>