export const getNowFormatDate = function (now) {
    let currentTimestamp = Date.now(); // 获取当前时间戳
    if (now) {
        const myDate = new Date(now);
        currentTimestamp = myDate.getTime();
    }
    // 创建一个表示时间戳的 Date 对象
    const currentDate = new Date(currentTimestamp);
    // 提取年、月、日
    const year = currentDate.getFullYear();
    const month = currentDate.getMonth() + 1; // 月份是从 0 到 11，因此需要加 1
    const day = currentDate.getDate();
    // 格式化为字符串
    return `${year}-${month < 10 ? '0' + month : month}-${day < 10 ? '0' + day : day}`;
}

// 获取本月第一天表示 yyyy-MM-01
export const getFirstDayofMouthFormatDate = function (now) {
    let currentTimestamp = Date.now(); // 获取当前时间戳
    if (now) {
        const myDate = new Date(now);
        currentTimestamp = myDate.getTime();
    }
    // 创建一个表示时间戳的 Date 对象
    const currentDate = new Date(currentTimestamp);
    // 提取年、月、日
    const year = currentDate.getFullYear();
    const month = currentDate.getMonth() + 1; // 月份是从 0 到 11，因此需要加 1
    const day = 1;
    // 格式化为字符串
    return `${year}-${month < 10 ? '0' + month : month}-${day < 10 ? '0' + day : day}`;
}

// 获取本年第一天表示 yyyy-1-1
export const getFirstDayofYearFormatDate = function (now) {
    let currentTimestamp = Date.now(); // 获取当前时间戳
    if (now) {
        const myDate = new Date(now);
        currentTimestamp = myDate.getTime();
    }
    // 创建一个表示时间戳的 Date 对象
    const currentDate = new Date(currentTimestamp);
    // 提取年、月、日
    const year = currentDate.getFullYear();
    const month = 1; // 月份是从 0 到 11，因此需要加 1
    const day = 1;
    // 格式化为字符串
    return `${year}-${month < 10 ? '0' + month : month}-${day < 10 ? '0' + day : day}`;
}
export const getSevenDaysAgoDate = function (now) {
    let currentTimestamp = Date.now(); // 获取当前时间戳
    if (now) {
        const myDate = new Date(now);
        currentTimestamp = myDate.getTime();
    }
    const sevenDaysAgoTimestamp = currentTimestamp - (7 * 24 * 60 * 60 * 1000);
    // 创建表示 7 天前的 Date 对象
    const sevenDaysAgoDate = new Date(sevenDaysAgoTimestamp);
    // 提取年、月、日
    const sevenDaysAgoYear = sevenDaysAgoDate.getFullYear();
    const sevenDaysAgoMonth = sevenDaysAgoDate.getMonth() + 1;
    const sevenDaysAgoDay = sevenDaysAgoDate.getDate();
    // 格式化为字符串
    return `${sevenDaysAgoYear}-${sevenDaysAgoMonth < 10 ? '0' + sevenDaysAgoMonth : sevenDaysAgoMonth}-${sevenDaysAgoDay < 10 ? '0' + sevenDaysAgoDay : sevenDaysAgoDay}`;
}

export const getOneMonthAgoDate = function (now) {
    let currentTimestamp = Date.now(); // 获取当前时间戳
    if (now) {
        const myDate = new Date(now);
        currentTimestamp = myDate.getTime();
    }
    const oneMonthAgoTimestamp = currentTimestamp - (30 * 24 * 60 * 60 * 1000); // 这里简单地用 30 天来近似一个月
    // 创建表示一个月前的 Date 对象
    const oneMonthAgoDate = new Date(oneMonthAgoTimestamp);
    // 提取年、月、日
    const oneMonthAgoYear = oneMonthAgoDate.getFullYear();
    const oneMonthAgoMonth = oneMonthAgoDate.getMonth() + 1;
    const oneMonthAgoDay = oneMonthAgoDate.getDate();
    // 格式化为字符串
    return `${oneMonthAgoYear}-${oneMonthAgoMonth < 10 ? '0' + oneMonthAgoMonth : oneMonthAgoMonth}-${oneMonthAgoDay < 10 ? '0' + oneMonthAgoDay : oneMonthAgoDay}`;
}
