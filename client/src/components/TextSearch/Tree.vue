<template>
<div>
    <TreeItem v-for="(child, idx) in models" :key="idx" :ref="`child`" :props_model="child" @onSelectAnyChild="onSelectAnyChild" />
</div>
</template>

<script>
import TreeItem from '@/components/TextSearch/TreeItem'

export default {
  name: 'Tree',
  components: {
    TreeItem
  },
  computed: {
  },
  props: {
    props_models: {
      type: Array,
      default: () => []
    }
  },
  data () {
    return {
      models: this.props_models
    }
  },
  methods: {
    onSelectAnyChild (child) {
      if (this.models) {
        var len = this.models.length || 0
        for (let i = 0; i < len; i++) {
          this.$refs['child'][i].onRequestUnselected()
        }
      }
      child.selected = true
      console.log(child)
    },
    onOpenAll () {
      if (this.model.children) {
        var len = this.model.children.length || 0
        for (let i = 0; i < len; i++) {
          this.$refs['child'][i].onOpenAll()
        }
      }
    },
    onCloseAll () {
      if (this.model.children) {
        var len = this.model.children.length || 0
        for (let i = 0; i < len; i++) {
          this.$refs['child'][i].onCloseAll()
        }
      }
    },
    onSelectParent () {
    }
  }
}
</script>

<style scoped lang="scss">
@import '@/scss/common.scss';

</style>
