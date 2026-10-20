<script lang="ts">
  interface Props {
    description: string;
  }

  let { description }: Props = $props();
</script>

<div class="card-description">
  <p>{description}</p>
</div>

<style lang="less">
  @import "../less/variables.less";
  @import "../less/mixins.less";
  @import "../less/card-description.less";
</style>
