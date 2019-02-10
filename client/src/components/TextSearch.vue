<template>
<div id="content" class="outermost-box">
    <div class="left-control">
        <div class="left-control__wrapper">
            <div class="gmenus left-control__row">
                <div class="button">Sql検索</div>
            </div>
            <div class="search-form">
                <div class="left-control__row">
                    <form @submit.prevent="searchTextdata">
                    <input type="search" class="search__textbox" v-model="query"><button class="search__button" type="submit">Go</button>
                    </form>
                </div>
                <div>
                    <input type="checkbox" id="TextSearch__wordOnly" v-model="wordOnly">
                    <label for="TextSearch__wordOnly" class="checkbox-label">単語のみ</label>
                    <input type="checkbox" id="TextSearch__not-ignorecase" v-model="notIgnoreCase">
                    <label for="TextSearch__not-ignorecase" class="checkbox-label">大文字小文字を区別</label>
                    <input type="checkbox" id="TextSearch__regex" v-model="regex">
                    <label for="TextSearch__regex" class="checkbox-label">正規表現</label>
                    <input type="checkbox" id="TextSearch__bodyOnly" v-model="bodyOnly">
                    <label for="TextSearch__bodyOnly" class="checkbox-label">本文のみ検索</label>
                </div>
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
        <pre class="sqlbody">{{ detailText }}</pre>
    </div>
</div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'TestSearch',
  data: () => ({
    query: '',
    wordOnly: false,
    notIgnoreCase: false,
    regex: false,
    bodyOnly: false,
    entries: [],
    selectedIdxies: [],
    detailCaption: '',
    detailText: '',
    detailTextSepalator: '\r\n----------------------------------\r\n'
  }),
  mounted: function () {
    this.searchTextdata()
  },
  methods: {
    searchTextdata: function () {
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
          if (data.Error) {
            alert('テキストデータの検索に失敗しました')
            return
          }
          this.entries = data.datas
          this.selectedIdxies = response.count > 0 ? [0] : []
        })
    },
    reloadSelectedText: function () {
      let entries = []
      for (let i of this.selectedIdxies) {
        entries.push(this.entries[i])
      }
      this.detailCaption = entries.map(e => e.caption).join(', ')
      this.detailText = entries.map(e => e.body).join(this.detailTextSepalator)
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

.search__textbox {
    vertical-align: middle;
    font-size: 12px;
    width: 180px;
    padding-left: 5px;
    display: inline-box;
    background: $leftCtlBGColor;
    color: $leftCtlTextColor;
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

.search-list__wrapper {
    margin-top: 15px;
    min-height: 200px;
    height: calc(100% - 160px);
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
}
.sqlbody {
    margin-top: 15px;
    height: calc(100% - 75px);
    background: #ddd;
    color: #333;
    padding: 15px;
    overflow-y: scroll;
}
</style>
