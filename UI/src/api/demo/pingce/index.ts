import request from '/@/utils/request';
//对任职信息进行操作
export function getpingce(query?:Object) {
    console.log("函数查询测试：",query);
    return request({
        url: '/api/v1/demo/pingce/list',
        method: 'get',
        params:query
    })
}
//精确查找
export function getpingce1(query?:Object) {
    console.log("函数查询测试：",query);
    return request({
        url: '/api/v1/demo/pingce1/list',
        method: 'get',
        params:query
    })
}
export function addpingce(data:object) {
    console.log("函数增加测试接口",data);
    return request({
        url: '/api/v1/demo/pingce/add',
        method: 'post',
        data:data
    })
}


export function editpingce(data:object) {
    console.log("函数修改测试：",data);
    return request({
        url: '/api/v1/demo/pingce/edit',
        method: 'put',
        data:data
    })
}


export function deletepingce(data:object) {
    console.log("函数删除测试：",data);
    return request({
        url: '/api/v1/demo/pingce/delete',
        method: 'delete',
        data:data
    })
}
