<template>
<div id="content" class="outermost-box">
    <div class="left-control">
        <div class="left-control__wrapper">
            <div class="left-control__header left-control__row">
                <div class="left-control__caption">Sql検索</div>
                <div class="left-control__config"><span class="config-icon" @click="$refs.SettingsDlg.show()"></span></div>
            </div>
            <div class="search-form">
                <div class="left-control__row">
                    <form @submit.prevent="searchTextdata">
                        <input type="search" class="search__textbox" v-bind:class="{ error: hasQueryAlert }" v-model="query" @keyup="refleshQueryAlert"><button class="search__button" type="submit">Go</button>
                    </form>
                </div>
                <div>
                    <input type="checkbox" id="TextSearch__wordOnly" v-model="wordOnly" @change="refleshQueryAlert">
                    <label for="TextSearch__wordOnly" class="checkbox-label">単語のみ</label>
                    <input type="checkbox" id="TextSearch__not-ignorecase" v-model="notIgnoreCase" @change="refleshQueryAlert">
                    <label for="TextSearch__not-ignorecase" class="checkbox-label">大文字小文字を区別</label>
                    <input type="checkbox" id="TextSearch__regex" v-model="regex" @change="refleshQueryAlert">
                    <label for="TextSearch__regex" class="checkbox-label">正規表現</label>
                    <input type="checkbox" id="TextSearch__bodyOnly" v-model="bodyOnly" @change="refleshQueryAlert">
                    <label for="TextSearch__bodyOnly" class="checkbox-label">本文のみ検索</label>
                </div>
                <div v-if="errorMessage !== ''" class="error-box">{{ errorMessage }}</div>
            </div>
            <div class="search-list__menutype-wrapper">
                <div @click="listmode = 'list'" class="search-list__icon-box" :class="{'selected': listmode === 'list'}"><img src="@/assets/icon-listmenu.png" class="search-list__icon"></div>
                <div @click="listmode = 'tree'" class="search-list__icon-box" :class="{'selected': listmode === 'tree'}"><img src="@/assets/icon-treemenu.png" class="search-list__icon treemenu"></div>
            </div>
            <div class="search-list__wrapper">
                <select v-model="selectedIdxies" multiple class="search-list" @change="reloadSelectedText">
                    <option v-for="entry in entries" v-bind:key="entry.idx" :value="entry.idx">{{entry.caption}}</option>
                </select>
            </div>
        </div>
    </div>
    <div class="detail-view">
        <h3 class="detail-caption">{{ detailCaption || '...' }}</h3>
        <div v-highlightjs="detailText" class="sqlbody"><code class="sql"></code></div>
    </div>
    <SettingsDlg ref="SettingsDlg"></SettingsDlg>
</div>
</template>

<script>
import axios from 'axios'
import SettingsDlg from './SettingsDlg'
import isvalid from '@/js/isvalid'

export default {
  name: 'TestSearch',
  components: {
    SettingsDlg
  },
  data: () => ({
    query: '',
    wordOnly: false,
    notIgnoreCase: false,
    regex: false,
    bodyOnly: false,
    entries: [],
    listmode: 'tree',
    selectedIdxies: [],
    detailCaption: '',
    detailText: '',
    detailTextSepalator: '\r\n----------------------------------\r\n',
    errorMessage: '',
    hasQueryAlert: false
  }),
  mounted: function () {
    this.searchTextdata()
  },
  methods: {
    searchTextdata: function () {
      this.setError('')
      const baseurl = this.$ownapi.resolveurl('/searchtext')
      const queryparm = [
        `query=${encodeURIComponent(this.query)}`,
        `wordOnly=${this.wordOnly}`,
        `ignoreCase=${!this.notIgnoreCase}`,
        `regex=${this.regex}`,
        `bodyOnly=${this.bodyOnly}`,
        'responseBody=true'
      ].join('&')

      axios.get(`${baseurl}?${queryparm}`)
        .then(response => {
          let data = response.data
          this.entries = data.datas
          this.selectedIdxies = response.count > 0 ? [0] : []
        })
        .catch(response => {
          this.setError('検索処理でエラーが発生しました。開発者向けヒント：F12開発者コンソールをご確認ください。')
          console.error(response)
        })
    },
    reloadSelectedText: function () {
      let entries = []
      for (let i of this.selectedIdxies) {
        entries.push(this.entries[i])
      }
      this.detailCaption = entries.map(e => e.caption).join(', ')
      this.detailText = entries.map(e => e.body).join(this.detailTextSepalator)
    },
    refleshQueryAlert: function () {
      this.hasQueryAlert = !this.checkQueryCond()
    },
    checkQueryCond: function () {
      if (this.regex && !isvalid.regex(this.query)) {
        return false
      }
      return true
    },
    setError: function (msg) {
      msg = msg || ''
      this.errorMessage = msg
    }
  }
}
</script>

<style scoped lang="scss">
@import '@/scss/common.scss';

::-webkit-scrollbar {
    width: 10px;
    height: 10px;
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

.outermost-box {
    @include grow-box;
    display: flex;
    display: -webkit-flex;
    flex-direction: row;
}

.left-control {
    @include grow-box;

    width: 260px;
    max-width: 260px;
    border-right: solid 1px $leftBGColorDark;

    background: $leftBGColor;
    color: $leftTextColor;

    & .disable {
        color: $leftTextColorDisable;
    }
}
.left-control__wrapper {
    @include grow-box;
    padding-top: 20px;
    padding-left: 18px;
}
.left-control__row {
    margin-top: 6px;
    height: 24px;
    line-height: 24px;
    margin-bottom: 6px;
}

.left-control__header {
    display: flex;
    flex-direction: row;
}
.left-control__caption {
    display: inline-block;
    flex-grow: 1;
}
.left-control__config {
    display: inline-block;
    margin-right: 20px;
}
.config-icon {
    display: inline-block;
    height: 22px;
    width: 22px;
    background: url('../assets/gear.png');
    background-size: 14px 14px;
    background-position: center;
    background-repeat: no-repeat;
    cursor: pointer;
    border-radius: 9px;

    &:hover {
        background-color: #00000088;
    }
}

.search__textbox {
    vertical-align: middle;
    font-size: 12px;
    width: 180px;
    padding-left: 5px;
    display: inline-box;
    background: $leftCtlBGColor;
    color: $leftCtlTextColor;
    &.error {
        color: #f82266;
        background-color: #fee8f0;
    }
}
.search__button {
    vertical-align: middle;
    height: 24px;
    display: inline-box;
    color: $leftTextColor;
    background: $leftPrimaryColor;
    width: 40px;
}

input[type=checkbox] {
  display: none;
}
.checkbox-label {
  font-size: 12px;
  box-sizing: border-box;
  -webkit-transition: background-color 0.2s linear;
  transition: background-color 0.2s linear;
  position: relative;
  display: inline-block;
  padding-top: 2px;
  padding-right: 5px;
  padding-left: 25px;
  border-radius: 8px;
  vertical-align: middle;
  cursor: pointer;
}

.checkbox-label:after {
  border-color: $leftTextColor;
}
.checkbox-label:after {
  -webkit-transition: border-color 0.1s linear;
  transition: border-color 0.1s linear;
  position: absolute;
  top: 60%;
  left: 3px;
  display: block;
  margin-top: -10px;
  width: 16px;
  height: 16px;
  border: 2px solid #bbb;
  border-radius: 6px;
  content: '';
}

.checkbox-label:before {
  -webkit-transition: opacity 0.1s linear;
  transition: opacity 0.1s linear;
  position: absolute;
  top: 55%;
  left: 9px;
  display: block;
  margin-top: -7px;
  width: 5px;
  height: 9px;
  border-right: 3px solid $leftTextColor;
  border-bottom: 3px solid $leftTextColor;
  content: '';
  opacity: 0;
  -webkit-transform: rotate(45deg);
  -ms-transform: rotate(45deg);
  transform: rotate(45deg);
}
input[type=checkbox]:checked + .checkbox-label:before {
  opacity: 1;
}

.error-box {
    color: #dc143c;
    font-size: 10px;
}

.search-list__menutype-wrapper {
  margin-top: 12px;
}
.search-list__icon-box {
  display: inline-block;
  width: 24px;
  height: 24px;
  cursor: pointer;
  border-radius: 2px;
  &:hover {
    background: rgba(0, 0, 0, 0.3);
  }
  &.selected, &:active {
    box-shadow: 10px 10px 10px 10px rgba(0,0,0,0.5) inset;
    background: rgba(0, 0, 0, 0.3);
  }
}

.search-list__icon {
  width: 12px;
  height: 12px;
  margin: 6px;
}

.search-list__wrapper {
    min-height: 200px;
    height: calc(100% - 190px);
}

.search-list {
    @include grow-box;
    width: 220px;
    height: 100%;
    background: $leftCtlBGColor;
    color: $leftCtlTextColor;

    option {
        padding-left: 5px;
        font-size: 11px;
    }
}

.detail-view {
    width: calc(100% - 300px);
    display: flex;
    display: -webkit-flex;
    flex-direction: column;
    flex-grow: 1;
    padding-top: 60px;
    padding-left: 25px;
    padding-right: 25px;
}

.detail-caption {
    font-size: 24px;
    height: 38px;
    overflow: hidden;
    text-overflow: ellipsis;
    word-wrap: break-word;
    white-space: nowrap;
}

.sqlbody {
    display: block;
    white-space: pre-wrap;
    word-wrap: break-word;
    height: calc(100% - 75px);
    padding: 15px;
    overflow-y: scroll;
    background: #dddde3;
}
</style>
<style>
.hljs, .hljs-subst {
    color: #444;
}
.hljs-comment {
    color: #888888;
}

.hljs-keyword,
.hljs-attribute,
.hljs-selector-tag,
.hljs-meta-keyword,
.hljs-doctag,
.hljs-name {
    font-weight: bold;
}

/* User color: hue: 0 */

.hljs-type,
.hljs-string,
.hljs-number,
.hljs-selector-id,
.hljs-selector-class,
.hljs-quote,
.hljs-template-tag,
.hljs-deletion {
    color: #880000;
}

.hljs-title,
.hljs-section {
    color: #880000;
    font-weight: bold;
}

.hljs-regexp,
.hljs-symbol,
.hljs-variable,
.hljs-template-variable,
.hljs-link,
.hljs-selector-attr,
.hljs-selector-pseudo {
    color: #BC6060;
}

/* Language color: hue: 90; */

.hljs-literal {
    color: #78A960;
}

.hljs-built_in,
.hljs-bullet,
.hljs-code,
.hljs-addition {
    color: #397300;
}

/* Meta color: hue: 200 */

.hljs-meta {
    color: #1f7199;
}

.hljs-meta-string {
    color: #4d99bf;
}

/* Misc effects */

.hljs-emphasis {
    font-style: italic;
}

.hljs-strong {
    font-weight: bold;
}
</style>
