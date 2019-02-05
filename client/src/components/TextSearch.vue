<template>
<div id="content" class="outermost-box">
    <div class="left-control">
        <div class="left-control__wrapper">
            <div class="gmenus left-control__row">
                <div class="button">Sql検索</div>
            </div>
            <div class="search-form">
                <div class="left-control__row">
                    <input type="search" class="search__textbox"><button class="search__button">Go</button>
                </div>
                <div>
                    <input type="checkbox" value="only-word" id="TextSearch__onlyword">
                    <lable class="checkbox-label" @click="toggleCheck('TextSearch__onlyword')">単語のみ</lable>
                    <input type="checkbox" value="only-word" id="TextSearch__not-ignorecase">
                    <lable class="checkbox-label" @click="toggleCheck('TextSearch__not-ignorecase')">大文字小文字を区別</lable>
                    <input type="checkbox" value="only-word" id="TextSearch__use-regex">
                    <lable class="checkbox-label" @click="toggleCheck('TextSearch__use-regex')">正規表現</lable>
                    <input type="checkbox" value="only-word" id="TextSearch__only-body">
                    <lable class="checkbox-label" @click="toggleCheck('TextSearch__only-body')">本文のみ検索</lable>
                </div>
            </div>
            <div class="search-list__wrapper">
                <select multiple class="search-list">
                    <option>a</option>
                    <option>b</option>
                </select>
            </div>
        </div>
    </div>
    <div class="detail-view">
        <h3 class="detail-caption">Title</h3>
        <div class="sqlbody">
            SELECT * FROM scheme.table;
        </div>
    </div>
</div>
</template>

<script>
export default {
  name: 'TestSearch',
  methods: {
    toggleCheck: function (id) {
      var el = document.querySelector('#' + id)
      if (el) {
        el.checked = !el.checked
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

.outermost-box {
    @include grow-box;
    display: flex;
    display: -webkit-flex;
    flex-direction: row;
}

.left-control {
    @include grow-box;

    min-width: 1px;
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
