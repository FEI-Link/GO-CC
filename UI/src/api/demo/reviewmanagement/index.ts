import request from '/@/utils/request';
//对任职信息进行操作
export function getreviewmanagement(query?:Object) {
    console.log("函数查询测试：",query);
    return request({
        url: '/api/v1/demo/reviewmanagement/list',
        method: 'get',
        params:query
    })
}

export function addreviewmanagement(data:object) {
    console.log("函数增加测试接口",data);
    return request({
        url: '/api/v1/demo/reviewmanagement/add',
        method: 'post',
        data:data
    })
}


export function editreviewmanagement(data:object) {
    console.log("函数修改测试：",data);
    return request({
        url: '/api/v1/demo/reviewmanagement/edit',
        method: 'put',
        data:data
    })
}


export function deletereviewmanagement(data:object) {
    console.log("函数删除测试：",data);
    return request({
        url: '/api/v1/demo/reviewmanagement/delete',
        method: 'delete',
        data:data
    })
}
