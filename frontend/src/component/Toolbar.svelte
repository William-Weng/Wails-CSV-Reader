<script lang="ts">
  let { totalCount, currentFilePath } = $props();
</script>

<header class="toolbar">
  <div class="toolbar-info">
    <h1>CSV 讀取器</h1>
    <span>共 {totalCount} 筆</span>
  </div>

  <div class="toolbar-file">
    {#if currentFilePath}
      <span class="file-label">目前檔案</span>
      <span class="file-path" title={currentFilePath}>{currentFilePath}</span>
    {:else}
      <span class="file-placeholder">請拖入 CSV 檔案</span>
    {/if}
  </div>
</header>

<style lang="less">
  @import "../less/variables.less";
  @import "../less/mixins.less";
  @import "../less/toolbar.less";
</style>
