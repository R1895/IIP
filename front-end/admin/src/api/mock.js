// fake api file, only for template use.
// data storaged in browser.
// Ziyu Niu 2023/12/4

import request from '@/utils/request'
import http from '@/utils/request'

// params.key and value
export function searchTables(params) {
    return new Promise((resolve,reject)=>{
        let dataStr = localStorage.getItem("mock")
        let data
        if(dataStr==null||dataStr==void 0 ||dataStr==""){
            data = []
        } else {
            data = JSON.parse(dataStr)
        }
        for(let i=0;i<data.length;i++){
            
        }
        localStorage.setItem("mock",JSON.stringify(data))
        resolve({
            Status:200,
            Data:[],
            Msg:"ok"
        })
    })
}

// getTables
export function getTables() {
    //
    return new Promise((resolve,reject)=>{
        let dataStr = localStorage.getItem("mock")
        let data
        if(dataStr==null||dataStr==void 0 ||dataStr==""){
            // init data
            let data = [
                {id:"0",title:"a",author:"a","pageviews":100,status:0,displayTime:new Date()},
                {id:"1",title:"b",author:"b","pageviews":101,status:1,displayTime:new Date()},
                {id:"2",title:"c",author:"c","pageviews":102,status:2,displayTime:new Date()}
            ]
            localStorage.setItem("mock",JSON.stringify(data))
        } else {
            data = JSON.parse(dataStr)
        }
        resolve({
            Status:200,
            Data:data,
            Msg:"ok"
        })
    })
}   

export function addTable(col) {
    return new Promise((resolve,reject)=>{
        let dataStr = localStorage.getItem("mock")
        let data
        if(dataStr==null||dataStr==void 0 ||dataStr==""){
            data = []
        } else {
            data = JSON.parse(dataStr)
        }
        data.push(col)
        localStorage.setItem("mock",JSON.stringify(data))
        resolve({
            Status:200,
            Data:[],
            Msg:"ok"
        })
    })
}

export function updateTable(col) {
    return new Promise((resolve,reject)=>{
        let dataStr = localStorage.getItem("mock")
        let data
        if(dataStr==null||dataStr==void 0 ||dataStr==""){
            data = []
        } else {
            data = JSON.parse(dataStr)
        }

        for(let i=0;i<data.length;i++){
            if(data[i].id==col.id){
                data[i] = col
                break;
            }
        }

        localStorage.setItem("mock",JSON.stringify(data))
        resolve({
            Status:200,
            Data:[],
            Msg:"ok"
        })
    })
}

export function deleteTable(id) {
    return new Promise((resolve,reject)=>{
        let dataStr = localStorage.getItem("mock")
        let data
        if(dataStr==null||dataStr==void 0 ||dataStr==""){
            data = []
        } else {
            data = JSON.parse(dataStr)
        }

        for(let i=0;i<data.length;i++){
            if(data[i].id==id){
                data.splice(i,1)
                break;
            }
        }
        console.log(data)
        localStorage.setItem("mock",JSON.stringify(data))
        resolve({
            Status:200,
            Data:[],
            Msg:"ok"
        })
    })
}

export function deleteTables(idList) {
    return new Promise((resolve,reject)=>{
        let dataStr = localStorage.getItem("mock")
        let data
        if(dataStr==null||dataStr==void 0 ||dataStr==""){
            data = []
        } else {
            data = JSON.parse(dataStr)
        }
        idList.forEach(col=>{
            for(let i=0;i<data.length;i++){
                if(data[i].id==col.id){
                    data.splice(i,1)
                    break;
                }
            }
        })

        localStorage.setItem("mock",JSON.stringify(data))
        resolve({
            Status:200,
            Data:[],
            Msg:"ok"
        })
    })
}
