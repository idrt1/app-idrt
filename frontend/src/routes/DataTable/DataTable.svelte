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
    import ChevronDown from "lucide-svelte/icons/chevron-down";
    import Actions from "./Data-table-actions.svelte";
    import DataTableCheckbox from "./Data-table-checkbox.svelte";
    import * as Table from "$lib/components/ui/table/index.js";
    import { Button } from "$lib/components/ui/button/index.js";
    import * as DropdownMenu from "$lib/components/ui/dropdown-menu/index.js";
    import { cn } from "$lib/utils.js";
    import { Input } from "$lib/components/ui/input/index.js";
    import {client} from "$lib/wailsjs/go/models";

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
            accessor: "nom",
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
            header: "Status",
            accessor: "categorie",
            plugins: { sort: { disable: true }, filter: { exclude: true } }
        }),
        table.column({
            header: "Email",
            accessor: ({ adresseElectronique }) => adresseElectronique,
            cell: ({ value }) => value.toLowerCase(),
            plugins: {
                sort: {
                    disable: false // Activation du tri
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
        })
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

    const { hiddenColumnIds } = pluginStates.hide;
    const ids = flatColumns.map((c) => c.id);
    let hideForId = Object.fromEntries(ids.map((id) => [id, true]));

    $: $hiddenColumnIds = Object.entries(hideForId)
        .filter(([, hide]) => !hide)
        .map(([id]) => id);

    const { hasNextPage, hasPreviousPage, pageIndex } = pluginStates.page;
    const { filterValue } = pluginStates.filter;

    const { selectedDataIds } = pluginStates.select;

    const hideableCols = ["categorie", "adresseElectronique", "nom"];
</script>

<div class="w-full">
    <div class="flex items-center py-4">
        <Input class="max-w-sm" placeholder="Filtrer par mails..." type="text" bind:value={$filterValue}/>
        <DropdownMenu.Root>
            <DropdownMenu.Trigger asChild let:builder>
                <Button variant="outline" class="ml-auto" builders={[builder]}>Columns <ChevronDown class="ml-2 h-4 w-4" /></Button>
            </DropdownMenu.Trigger>
            <DropdownMenu.Content>
                {#each flatColumns as col}
                    {#if hideableCols.includes(col.id)}
                        <DropdownMenu.CheckboxItem bind:checked={hideForId[col.id]}>
                            {col.header}
                        </DropdownMenu.CheckboxItem>
                    {/if}
                {/each}
            </DropdownMenu.Content>
        </DropdownMenu.Root>
    </div>
    <div class="rounded-md border">
        <Table.Root {...$tableAttrs}>
            <Table.Header>
                {#each $headerRows as headerRow}
                    <Subscribe rowAttrs={headerRow.attrs()}>
                        <Table.Row>
                            {#each headerRow.cells as cell (cell.id)}
                                <Subscribe attrs={cell.attrs()} let:attrs props={cell.props()} let:props>
                                    <Table.Head {...attrs} class={cn("[&:has([role=checkbox])]:pl-3")}>
                                        {#if cell.id === "Nom"}
                                            <div class="text-right font-medium">
                                                <Render of={cell.render()} />
                                            </div>
                                        {:else if cell.id === "Email"}
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
                        <Table.Row {...rowAttrs} data-state={$selectedDataIds[row.id] && "selected"}>
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