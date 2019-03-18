<template>
<dialog id="settingsdlg" class="settings-wrapper">
  <header>
    <h3>設定</h3>
    <div class="close-button-wrapper">
      <div class="close-button" @click="close">×</div>
    </div>
  </header>
  <div class="content">
    <div v-if="textdata_msg != ''" class="error">{{ textdata_msg }}</div>
    <div v-if="currentTab === 'textdata'">
      <h4>検索データをアップロードする</h4>
      <div class="section">
        <input id="config_textdatafile" type="file" name="file" @change="selectTextdataFile">
        <button type="submit"
                v-bind:class="{ 'primary-button': textdata_okupload, 'disabled-button': !textdata_okupload }"
                @click="uploadTextdata">アップロード</button>
      </div>
      <h4>検索データ一覧</h4>
      <div class="section">
        <table>
          <tr>
            <th>ファイル</th>
            <th>更新日時</th>
            <th>アクション</th>
          </tr>
          <tr v-for="textdata in textdatas" v-bind:key="textdata.filename">
            <td>{{ textdata.filename }}</td>
            <td>{{ textdata.lastupdate }}</td>
            <td class="datafiles-col">
              <button @click="download(textdata.filename)"><img width="16px" height="16px" src="@/assets/download.png"></button>
              <button @click="remove(textdata.filename)"><img width="16px" height="16px" src="@/assets/garbage.png"></button>
            </td>
          </tr>
        </table>
      </div>
    </div>
  </div>
</dialog>
</template>

<script>
import axios from 'axios'

export default {
  name: 'SettingsDlg',
  data: () => ({
    currentTab: 'textdata',
    textdatas: [],
    textdata_uploadfile: null,
    textdata_msg: ''
  }),
  computed: {
    textdata_okupload: function () {
      return this.textdata_uploadfile !== null
    }
  },
  methods: {
    show: function () {
      let dlg = document.querySelector('#settingsdlg')
      if (dlg && !dlg.open) {
        dlg.showModal()
        this.resetForm()
        this.requestReloadTextdata()
      }
    },
    close: function () {
      let dlg = document.querySelector('#settingsdlg')
      if (dlg && dlg.open) {
        dlg.close()
      }
    },
    resetForm: function () {
      document.querySelector('#config_textdatafile').value = ''
      this.textdata_uploadfile = null
      this.textdatas = []
    },
    requestReloadTextdata: function () {
      const baseurl = this.$ownapi.resolveurl('/datafile/list')
      const queryparm = [].join('&')
      this.textdata_msg = ''

      axios.get(`${baseurl}?${queryparm}`)
        .then(response => {
          let data = response.data
          this.textdatas = data.datas
        })
        .catch(response => {
          this.textdata_msg = 'テキストデータ一覧が取得できませんでした。開発者向けヒント：F12開発者コンソールをご確認ください。'
          console.error(response)
        })
    },
    uploadTextdata: function () {
      if (!this.textdata_uploadfile) {
        return
      }
      const baseurl = this.$ownapi.resolveurl('/datafile/upload')

      let formData = new FormData()
      formData.append('file', this.textdata_uploadfile)

      const config = {
        headers: {
          'content-type': 'multipart/form-data'
        }
      }

      axios.post(`${baseurl}`, formData, config)
        .then(() => {
          this.resetForm()
          this.requestReloadTextdata()
          this.textdata_msg = 'ファイルをアップロードしました'
        })
        .catch(response => {
          this.textdata_msg = 'アップロードが失敗しました。開発者向けヒント：F12開発者コンソールをご確認ください。'
          console.error(response)
        })
    },
    selectTextdataFile: function (e) {
      e.preventDefault()
      let files = e.target.files
      this.textdata_uploadfile = files[0]
    },
    download: function (filename) {
      const baseurl = this.$ownapi.resolveurl(`/datafile/download/${filename}`)
      axios.get(baseurl)
        .then((response) => {
          console.log(response)
          const blob = new Blob([response.data], { type: response.headers['Content-Type'] })
          let link = document.createElement('a')
          link.href = window.URL.createObjectURL(blob)
          link.download = filename
          link.click()
        })
        .catch((response) => {
          this.textdata_msg = `${filename}のダウンロードに失敗しました。開発者向けヒント：F12開発者コンソールをご確認ください。`
          console.error(response)
        })
    },
    remove: function (filename) {
      if (!confirm(`${filename}を削除しますか？`)) {
        return
      }

      const baseurl = this.$ownapi.resolveurl(`/datafile/remove/${filename}`)
      axios.delete(`${baseurl}`)
        .then(() => {
          this.resetForm()
          this.requestReloadTextdata()
          this.textdata_msg = `${filename}を削除しました`
        })
        .catch(response => {
          this.textdata_msg = `${filename}の削除に失敗しました。開発者向けヒント：F12開発者コンソールをご確認ください。`
          console.error(response)
        })
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
  width: 600px;
  border: none;
  box-shadow: 0px 0px 5px #000000cc;
  margin: 0px auto;
  height: 70%;
  min-height: 150px;
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
  padding: 15px 24px;
  font-size: 12px;
  overflow: scroll-y;
  height:100%;
}

h4 {
  font-size: 16px;
  margin-top: 5px;
  margin-bottom: 8px;
}

.section {
  margin-bottom: 36px;
}

table {
  border-collapse: collapse;
}

th, td {
  border: solid 1px $grayLine;
  padding: 3px 5px;
}

.datafiles-col {
  text-align: center;
  letter-spacing: 3px;
}

.error {
  margin-bottom: 16px;
}
</style>
