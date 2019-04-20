<template>
<div>
    <div class="treeitem-line" @click="select" :class="{'selected': selected, 'toggleable': !selectable}">
        <span class="treeitem-mark">{{ mark }}</span>
        <span class="treeitem-label">{{ label }}</span>
    </div>
    <div class="treeitem-children" :class="{ 'closed' : childClosed }">
        <div v-for="(child, idx) in model.children" :key="idx" >
            <TreeItem ref="child" :props_model="child" @onSelectAnyChild="onSelectAnyChild" />
        </div>
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
        label += `(${this.childrenCount})`
      }
      return label
    },
    childrenCount () {
      const getcnt = (model) => {
        if (model.children) {
          let sum = 0
          for (let m of model.children) {
            sum += getcnt(m)
          }
          return sum
        }
        return 1
      }

      if (this.model.children) {
        return getcnt(this.model)
      } else {
        return 1
      }
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
  height: 20px;
  &:hover {
    background: $primaryColor;
  }
  &.selected {
    background: $primaryColorLight;
    color: $primaryTextColor;
  }
  cursor: pointer;

  &.toggleable {
    background: rgba(0, 0, 0, 0.2);
    &:hover {
      background: $primaryColor;
    }
  }
}

.treeitem-mark {
  display: inline-block;
  width: 14px;
}

.treeitem-label {
  display: inline-block;
  flex-grow: 1;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.treeitem-children {
  padding-left: 14px;
  &.closed {
    display: none;
  }
  overflow: hidden;
}
</style>
