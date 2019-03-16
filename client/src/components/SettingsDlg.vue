<template>
<dialog id="settingsdlg" class="settings-wrapper">
  <header>
    <h3>コントロール</h3>
    <div class="close-button-wrapper">
      <div class="close-button" @click="close">×</div>
    </div>
  </header>
  <div class="content">
    
  </div>
</dialog>
</template>

<script>
export default {
  name: 'SettingsDlg',
  data: () => ({
  }),
  methods: {
    show: function () {
      let dlg = document.querySelector('#settingsdlg')
      if (dlg && !dlg.open) {
        dlg.showModal()
      }
    },
    close: function () {
      let dlg = document.querySelector('#settingsdlg')
      if (dlg && dlg.open) {
        dlg.close()
      }
    }
  }
}
</script>

<style scoped lang="scss">
@import '@/scss/common.scss';

::-webkit-scrollbar {
    width: 10px;
}

/*スクロールバーの軌道*/
::-webkit-scrollbar-track {
  background-color: rgba(0, 0, 0, 0.1);
}

/*スクロールバーの動く部分*/
::-webkit-scrollbar-thumb {
  background-color: rgba(0, 0, 50, .5);
  border-radius: 10px;
  box-shadow:0 0 0 1px rgba(255, 255, 255, .3);
}

dialog {
  min-width: 600px;
  border: none;
  box-shadow: 0px 0px 5px #000000cc;
  margin: 0px auto;
}

dialog::backdrop {
  position: fixed;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  background: #000000aa;
  width: 100%;
  height: 100%;
}

header {
  background: $leftBGColor;
  color: $leftTextColor;

  display: flex;
  display: -webkit-flex;
  flex-direction: row;

  padding: 6px 10px;
}

h3 {
  flex-grow: 1;
}

.close-button-wrapper {
  vertical-align: middle;
}

.close-button {
  display: inline-box;
  font-size: 18px;
  width: 24px;
  height: 24px;
  text-align: center;
  cursor: pointer;

  &:hover {
    background-color: #ffffff55;
  }
}

.content {
  padding: 10px;
}

</style>
