<template>
<div>
    <div class="treeitem-line" @click="select" :class="{'selected': selected}">
        <span class="treeitem-mark">{{ mark }}</span>
        <span class="treeitem-label">{{ label }}</span>
    </div>
    <div class="treeitem-children" :class="{ 'closed' : childClosed }">
        <TreeItem v-for="(child, idx) in model.children" :key="idx" :ref="'child'" :props_model="child" @onSelectAnyChild="onSelectAnyChild" />
    </div>
</div>
</template>

<script>
export default {
  name: 'TreeItem',
  components: {
    TreeFolderContents: () => import('@/components/TextSearch/TreeItem')
  },
  computed: {
    mark () {
      if (!this.model.children) {
        return ''
      }
      return this.childClosed ? '+' : '-'
    },
    label () {
      let label = this.model.label || 'item'
      if (this.model.children) {
        let len = this.model.children.length
        if (len) {
          label += `${len}`
        }
      }
      return label
    },
    selectable () {
      return !this.model.children
    }
  },
  data () {
    return {
      model: this.props_model,
      childClosed: this.props_childClosed,
      selected: this.props_selected
    }
  },
  props: {
    props_model: {
      type: Object,
      default: () => ({})
    },
    props_childClosed: {
      type: Boolean,
      default: () => false
    },
    props_selected: {
      type: Object,
      default: () => undefined
    }
  },
  methods: {
    select () {
      if (this.selectable) {
        this.$emit('onSelectAnyChild', this)
      } else {
        this.childClosed = !this.childClosed
      }
    },
    onSelectAnyChild (value) {
      this.$emit('onSelectAnyChild', value)
    },
    onRequestUnselected () {
      if (this.model.children) {
        var len = this.model.children.length || 0
        for (let i = 0; i < len; i++) {
          this.$refs['child'][i].onRequestUnselected(this.model)
        }
      }
      this.selected = false
    },
    onOpenAll () {
      if (this.model.children) {
        var len = this.model.children.length || 0
        for (let i = 0; i < len; i++) {
          this.$refs['child'][i].onOpenAll()
        }
      }
      this.childClosed = false
    },
    onCloseAll () {
      if (this.model.children) {
        var len = this.model.children.length || 0
        for (let i = 0; i < len; i++) {
          this.$refs['child'][i].onCloseAll()
        }
      }
      this.childClosed = true
    }
  }
}
</script>

<style scoped lang="scss">
@import '@/scss/common.scss';

.treeitem-line {
  padding-left: 3px;
  display: flex;
  &:hover {
    background: $primaryColor;
  }
  &.selected {
    background: $primaryColorLight;
    color: $primaryTextColor;
  }
  cursor: pointer;
}

.treeitem-mark {
  display: inline-block;
  width: 14px;
}

.treeitem-label {
  display: inline-block;
  flex-grow: 1;
}

.treeitem-children {
  padding-left: 14px;
  &.closed {
    display: none;
  }
  overflow: hidden;
}
</style>
