import request from '../until/request'
export function getScreenList(){
    return request.get('/api/tv/list')
}

export function changeAllStatus(showing) {
    return request.post('/api/tv/isplay/',{
        "data":showing
    })
}

export function getTvShow(terms) {
    let result = []
    for (let i = 0;i < terms.length;i++) {
        request.get('http://rap2api.taobao.org/app/mock/279014/api/tv/query?ip=' + terms[i].ip + '&id=' + terms[i].id).then(function (res) {
            result.push({
                tvIp:terms[i].ip,
                data:res.data
            })
        })
    }
    return result
}

export function getTvShowing(terms, program) {
    let result = []
    for (let i = 0;i < program;i++) {
        if (program[i].data.status === "play") {
            result.push({
                "ip": program[i].tvIp,
                "id": program[i].data.id,
                "currentShowName": program[i].data.curMaterialName,
            })
        }
    }
    return result
}

export function tvOperate(showID, id, ip, operation, offset) {
    return request.post('/api/tv/operate', {
        "Showid": showID,
        "id":id,
        "ip":ip,
        "operation":operation,
        "offset":offset,
    })
}

