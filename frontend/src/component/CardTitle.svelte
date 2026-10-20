<script lang="ts">
    interface Props { title: string; url?: string | null; openExternalURL: (url: string) => void | Promise<void>; }
    let { title, url, openExternalURL }: Props = $props();
</script>

{#if url}
    <button class="card-title clickable" type="button" {title} onclick={ () => { if (url) { void openExternalURL(url); }} }>
        <span>{title}</span>
    </button>
{:else}
    <h2 class="card-title" {title}>
        <span>{title}</span>
    </h2>
{/if}

<style lang="less">
  @import "../less/variables.less";
  @import "../less/mixins.less";
  @import "../less/card-title.less";
</style>
