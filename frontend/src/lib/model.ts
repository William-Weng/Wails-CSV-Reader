/**
 * 支援的原生系統對話框類型
 * - info：一般資訊或操作成功
 * - error：操作失敗或發生錯誤
 * - warning：提醒使用者注意或補齊必要輸入
 */
type DialogType = "info" | "error" | "warning";

type DroppedData = {
  elementID: string;
  path: string;
};

type ResourceItem = {
	title: string;
	description: string;
	url: string;
	example: string;
	types: string[];
};

type ResourceListResult = {
	types: string[];
	resources: ResourceItem[];
};

type CardItem = {
    id: string;
    title: string;
    description: string;
    url?: string;
    types: string[];
    example?: string;
};

type TenColors = readonly [
    string,
    string,
    string,
    string,
    string,
    string,
    string,
    string,
    string,
    string,
];
