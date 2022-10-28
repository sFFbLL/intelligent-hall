<template>
    <div class="main">
        <div class="one">
            <swt-btn partName="总控" class="allControl"></swt-btn>
        </div>
        <hr>
        <div class="two">
            <div class="terminal" @click="changeLocking(item.ip, item.id)" v-for="(item, i) in terms">0{{ i + 1 }}</div>
        </div>
        <div class="video">
            <p>正在播放：{{ locking.currentShowName }}</p>
            <div class="progressBar">
                <van-slider v-model="value" @change="onChange" bar-height="3px"
                            active-color="rgb(183,183,183)" inactive-color="white"/>
            </div>
            <div class="button">

            </div>
        </div>
    </div>
</template>

<script>
    import {changeAllStatus, getScreenList, getTvShow, getTvShowing, tvOperate} from '../api/screen'
    import SwtBtn from "../components/SwtBtn.vue";
    import {Toast} from 'vant';

    export default {
        name: "Screen",
        components: {SwtBtn},
        data() {
            return {
                //终端列表
                terms: [{ip: "1", id: 1}],
                showing: [
                    {
                        ip: "",
                        id: 0,
                        currentShowName:"",
                    }
                ],
                program: [],
                value: 50,
                locking: {
                    ip:0,
                    id:0,
                    showId:0,
                    currentShowName:"",
                },
            }
        },
        async mounted() {
            setInterval(() => {
                getScreenList().then(function (res) {
                    this.terms = res.data.termList.terms
                })
            }, 1000)
            setInterval(() => {
                this.program = getTvShow(this.terms)
            }, 1000)
            setInterval(() => {
                this.showing = getTvShowing(this.terms, this.program)
            }, 1000)
        },
        methods: {
            allControl() {
                changeAllStatus(this.showing)
            },
            tvOperation() {
                tvOperate()
            },
            onChange(value) {
                // Toast('当前值：' + value);
            },
            changeLocking(ip, id) {
                let showId = 0;
                let currentShowName = ""
                for (let i = 0;i < this.showing.length;i++) {
                    if (this.showing[i].ip === ip){
                        showId = this.showing[i].id
                        currentShowName = this.showing[i].currentShowName
                    }
                }
                this.locking = {
                    "ip": ip,
                    "id": id,
                    "showId": showId,
                    "currentShowName": currentShowName,
                }
            }
        }
    }
</script>

<style>
    * {
        margin: 0;
        padding: 0;
        border: 0;
    }

    body {
        background-color: rgb(230, 238, 243);
        height: 765px;
    }

    .main {
        display: flex;
        flex-direction: column;
        margin: auto auto;
        width: 65%;
        height: 500px;
        box-shadow: 5px 2px 10px rgba(0, 0, 0, 0.452), -2px -2px 10px white;
        border-radius: 25px;
    }

    .one {
        flex: 1;
        width: 580px;
        height: 70px;
        position: relative;
        margin: 0 0 30px;
    }

    .allControl .partName {
        color: #2b94ff;
    }

    .allControl .van-switch__node {
        background-color: #2b94ff;
    }

    hr {
        width: 97%;
        height: 2px;
        background-color: #ffffff;
        margin: 0 auto;
    }

    .two {
        display: flex;
        width: 580px;
        flex: 2;
        flex-wrap: wrap;
        overflow-y: scroll;
        margin: 30px auto 0;
    }

    .terminal {
        width: 75px;
        height: 50px;
        background-color: black;
        border: solid white 2px;
        color: white;
        margin-right: 17px;
        margin-bottom: 14px;
        text-align: center;
        line-height: 50px;
        box-shadow: 5px 2px 10px rgba(0, 0, 0, 0.452);
    }

    .terminal:hover {
        background-color: #2b94ff;
    }

    .video {
        display: flex;
        flex-direction: column;
        width: 580px;
        margin: 20px auto 20px;
        border: solid white 2px;
        border-radius: 10px;
        flex: 1.8;
    }

    .video p {
        text-align: center;
        margin: 25px;
        color: rgb(130, 129, 129);
    }

    .progressBar {
        width: 90%;
        margin: 0 auto;
    }

    .van-slider {
        width: 100%;
    }

    .button {
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    footer {
        text-align: center;
    }
</style>
