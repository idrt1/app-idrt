<script lang="ts">
    import {
        Render,
        Subscribe,
        createRender,
        createTable
    } from "svelte-headless-table";
    import {
        addHiddenColumns,
        addPagination,
        addSelectedRows,
        addSortBy,
        addTableFilter
    } from "svelte-headless-table/plugins";
    import { readable } from "svelte/store";
    import ArrowUpDown from "lucide-svelte/icons/arrow-up-down";
    import Actions from "./Data-table-actions.svelte";
    import DataTableCheckbox from "./Data-table-checkbox.svelte";
    import * as Table from "$lib/components/ui/table/index.js";
    import { Button } from "$lib/components/ui/button/index.js";
    import { cn } from "$lib/utils.js";
    import { Input } from "$lib/components/ui/input/index.js";
    import {client} from "$lib/wailsjs/go/models";
    import {goto} from "$app/navigation";
    import ArrowUp from "lucide-svelte/icons/arrow-up";

    export let clients: client.Client[] = [];


    const table = createTable(readable(clients), {
        sort: addSortBy({ disableMultiSort: true }),
        page: addPagination(),
        filter: addTableFilter({
            fn: ({ filterValue, value }) => value.includes(filterValue)
        }),
        select: addSelectedRows(),
        hide: addHiddenColumns()
    });

    const columns = table.createColumns([
        table.column({
            header: (_, { pluginStates }) => {
                const { allPageRowsSelected } = pluginStates.select;
                return createRender(DataTableCheckbox, {
                    checked: allPageRowsSelected
                });
            },
            accessor: "id",
            cell: ({ row }, { pluginStates }) => {
                const { getRowState } = pluginStates.select;
                const { isSelected } = getRowState(row);

                return createRender(DataTableCheckbox, {
                    checked: isSelected
                });
            },
            plugins: {
                sort: {
                    disable: true
                },
                filter: {
                    exclude: true
                }
            }
        }),
        table.column({
            header: "Prénom",
            accessor: ({ prenom }) => prenom,
            cell: ({ value }) => value.toLowerCase(),
            plugins: {
                sort: {
                    disable: false
                },
                filter: {
                    getFilterValue(value) {
                        return value.toLowerCase();
                    }
                }
            }
        }),
        table.column({
            header: "Nom",
            accessor: ({ nom }) => nom,
            cell: (item) => {
                return createRender(Actions, { nom: item.value });
            },
            plugins: {
                sort: {
                    disable: true
                }
            }
        }),
        table.column({
            header: "Email",
            accessor: ({ adresseElectronique }) => adresseElectronique,
            cell: ({ value }) => value.toLowerCase(),
            plugins: {
                sort: {
                    disable: false
                },
                filter: {
                    getFilterValue(value) {
                        return value.toLowerCase();
                    }
                }
            }
        }),
    ]);


    const {
        headerRows,
        pageRows,
        tableAttrs,
        tableBodyAttrs,
        flatColumns,
        pluginStates,
        rows
    } = table.createViewModel(columns);

    const { sortKeys } = pluginStates.sort;

    const { hasNextPage, hasPreviousPage, pageIndex } = pluginStates.page;
    const { filterValue } = pluginStates.filter;

    const { selectedDataIds } = pluginStates.select;

    function handleRowClick(clientId: number) {
        goto(`/clients?id=${clientId}`);
        console.log(clientId);
    }


    function exportData() {
        const csvContent = "data:text/csv;charset=utf-8,"
            + [
                "Nom",
                "Prénom",
                "Email",
                "Numéro de portable",
                "Catégorie",
                "Conjoint syndiqué",
                "Cotisation spéciale",
                "Cotisation",
                "Coût",
                "D ou N",
                "Date d'installation",
                "Date de création",
                "Date de paiement",
                "ID Syndiqué",
                "Numéro de téléphone professionnel",
                "Premier an cotisation",
                "Syndiqué",
                "Titre",
                "Type d'installation",
                "Catégorie cotisation CAT",
                "Catégorie cotisation DEP",
                "Adresse professionnelle 1",
                "Adresse professionnelle 2",
                "Âge",
                "Association",
                "Code postal professionnel",
                "Date de diplôme",
                "Date de modification",
                "Date de naissance",
                "Diplôme faculté",
                "Lieu d'exercice",
                "Numéro de téléphone domicile",
                "Pays professionnel",
                "Personne aide assistante",
                "Personne assistante",
                "Personne collaborateur",
                "Personne femme de ménage",
                "Personne laboratoire",
                "Personne réceptionniste",
                "Personnels",
                "Remarques",
                "Responsable",
                "Sexe",
                "Type d'exercice",
                "Ville professionnelle"
            ].join(";") + "\n"
            + clients.map(client => [
                client.nom,
                client.prenom,
                client.adresseElectronique,
                client.numeroPort,
                client.categorie,
                client.conjointSynd,
                client.cotiSpeciale,
                client.cotisation,
                client.cout,
                client.d_ou_N,
                client.dateInstallation,
                client.dateCreation,
                client.datePaiement,
                client.idSyndique,
                client.numeroTelProf,
                client.premierAnCoti,
                client.syndique,
                client.titre,
                client.typeInstallation,
                client.categorieCotiCat,
                client.categorieCotiDep,
                client.adresseProf1,
                client.adresseProf2,
                client.age,
                client.association,
                client.codePostalProf,
                client.dateDiplome,
                client.dateModification,
                client.dateNaissance,
                client.diplomeFaculte,
                client.lieuExercice,
                client.numeroTelDomicile,
                client.paysProf,
                client.persAideAssistante,
                client.persAssistante,
                client.persCollaborateur,
                client.persFemmeDeMenage,
                client.persLaboratoire,
                client.persReceptionniste,
                client.personnels,
                client.remarques,
                client.responsable,
                client.sexe,
                client.typeExercice,
                client.villeProf
            ].join(";")).join("\n");

        const encodedUri = encodeURI(csvContent);
        const link = document.createElement("a");
        link.setAttribute("href", encodedUri);
        link.setAttribute("download", "clients.csv");
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
    }



</script>

<div class="w-full pt-4 pl-8">
    <div class="flex items-center space-x-2 justify-between">
        <Input class="max-w-sm" placeholder="Rechercher..." type="text" bind:value={$filterValue}/>
        <Button variant="outline" on:click={exportData}>
            Exporter
            <ArrowUp class="ml-2 h-4 w-4"/>
        </Button>
    </div>
    <div class="rounded-md border mt-4">
        <Table.Root {...$tableAttrs}>
            <Table.Header>
                {#each $headerRows as headerRow}
                    <Subscribe rowAttrs={headerRow.attrs()}>
                        <Table.Row>
                            {#each headerRow.cells as cell (cell.id)}
                                <Subscribe attrs={cell.attrs()} let:attrs props={cell.props()} let:props>
                                    <Table.Head {...attrs} class={cn("[&:has([role=checkbox])]:pl-3")}>
                                        {#if cell.id === "Nom"}
                                            <div class="font-medium">
                                                <Render of={cell.render()} />
                                            </div>
                                        {:else if cell.id === "Prenom"}
                                            <Button variant="ghost" on:click={props.sort.toggle}>
                                                <Render of={cell.render()} />
                                                <ArrowUpDown class={cn($sortKeys[0]?.id === cell.id && "text-foreground","ml-2 h-4 w-4")}/>
                                            </Button>
                                        {:else}
                                            <Render of={cell.render()} />
                                        {/if}
                                    </Table.Head>
                                </Subscribe>
                            {/each}
                        </Table.Row>
                    </Subscribe>
                {/each}
            </Table.Header>

            <Table.Body {...$tableBodyAttrs}>
                {#each $pageRows as row (row.id)}
                    <Subscribe rowAttrs={row.attrs()} let:rowAttrs>
                        <Table.Row {...rowAttrs} data-id={row.id} data-state={$selectedDataIds[row.id] && "selected"} on:click={() => handleRowClick(row.original.id)}>
                            {#each row.cells as cell (cell.id)}
                                <Subscribe attrs={cell.attrs()} let:attrs>
                                    <Table.Cell class="[&:has([role=checkbox])]:pl-3" {...attrs}>
                                        {#if cell.id === "amount"}
                                            <div class="text-right font-medium">
                                                <Render of={cell.render()} />
                                            </div>
                                        {:else if cell.id === "status"}
                                            <div class="capitalize">
                                                <Render of={cell.render()} />
                                            </div>
                                        {:else}
                                            <Render of={cell.render()} />
                                        {/if}
                                    </Table.Cell>
                                </Subscribe>
                            {/each}
                        </Table.Row>
                    </Subscribe>
                {/each}
            </Table.Body>

        </Table.Root>
    </div>
    <div class="flex items-center justify-end space-x-2 py-4">
        <div class="text-muted-foreground flex-1 text-sm">
            {Object.keys($selectedDataIds).length} / {$rows.length} sélectionné(s)
        </div>
        <Button variant="outline" size="sm" on:click={() => ($pageIndex = $pageIndex - 1)} disabled={!$hasPreviousPage}>Précédent</Button>
        <Button variant="outline" size="sm" disabled={!$hasNextPage} on:click={() => ($pageIndex = $pageIndex + 1)}>Suivant</Button>
    </div>
</div>