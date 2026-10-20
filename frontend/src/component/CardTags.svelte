<script lang="ts">
  interface Props {
    types: string[];
    getTypeColor: (type: string) => string;
  }

  let { types, getTypeColor }: Props = $props();
</script>

{#if types.length > 0}
  <div class="card-tags" aria-label="分類">
    {#each types as type (type)}
      <span
        class="type-tag"
        style:--type-color={getTypeColor(type)}
      >
        {type}
      </span>
    {/each}
  </div>
{/if}

<style lang="less">
  @import "../less/variables.less";
  @import "../less/mixins.less";
  @import "../less/card-tags.less";
</style>
