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

    const df = new DateFormatter("en-US", { dateStyle: "long" });
    let startDate: DateValue | undefined = undefined;
    let endDate: DateValue | undefined = undefined;
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
            <Input id="syndicat-coti" placeholder="Entré un premier an" />
        </div>
        <div class="space-y-2">
            <Label for="syndicat-number">Union Number</Label>
            <Input id="syndicat-number" placeholder="Enter union number" />
        </div>
    </CardContent>
</Card>
