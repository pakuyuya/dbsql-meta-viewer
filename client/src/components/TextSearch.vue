<template>
<div id="content" class="outermost-box">
    <div class="left-control"></div>
    <div class="detail"></div>
</div>
</template>

<script>
export default {
  name: 'TestSearch'
}
</script>

<style scoped lang="scss">

.outermost-box {
    display: flex;
    display: -webkit-flex;
    flex-direction: row;
}

.left-control {
    display: flex;
    display: -webkit-flex;
    flex-direction: columns;
    flex-grow: 1;
}

.detail {
    display: flex;
    display: -webkit-flex;
    flex-direction: columns;
    flex-grow: 1;
}

</style>
