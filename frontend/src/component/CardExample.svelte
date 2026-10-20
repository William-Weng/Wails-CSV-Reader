<script lang="ts">
  interface Props {
    example?: string | null;
  }

  let { example }: Props = $props();
</script>

{#if example}
  <div class="card-example">
    <pre><code>{example}</code></pre>
  </div>
{/if}

<style lang="less">
  @import "../less/variables.less";
  @import "../less/mixins.less";
  @import "../less/card-example.less";
</style>
