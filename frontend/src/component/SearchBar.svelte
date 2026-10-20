<script lang="ts">
  let { searchText = $bindable(""), totalCount, filteredCount } = $props();
</script>

<footer class="search-bar">
  <div class="search-content">
    <label for="card-search">Search</label>
    <div class="search-input-wrap">
      <input id="card-search" type="search" placeholder="搜尋標題、說明、網址、範例或分類…" bind:value={searchText} onkeydown={(event) => { if (event.key === "Escape") { searchText = ""; }} }/>
    </div>
    <span class="search-count" role="status">{filteredCount} / {totalCount} 筆</span>
  </div>
</footer>

<style lang="less">
  @import "../less/variables.less";
  @import "../less/mixins.less";
  @import "../less/search.less";
</style>
