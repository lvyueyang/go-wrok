export const PERMISSION_CODE = {
  /** 查询C端用户列表 */
  ADMIN_C_USER_FIND_LIST: 'admin:c_user:find:list',
  /** 修改C端用户状态 */
  ADMIN_C_USER_UPDATE_STATUS: 'admin:c_user:update:status',
  /** 查询管理员列表 */
  ADMIN_USER_FIND_LIST: 'admin:user:find:list',
  /** 创建管理员 */
  ADMIN_USER_CREATE: 'admin:user:create',
  /** 修改管理员基本信息 */
  ADMIN_USER_UPDATE_INFO: 'admin:user:update:info',
  /** 删除管理员 */
  ADMIN_USER_DELETE: 'admin:user:delete',
  /** 修改管理员密码 */
  ADMIN_USER_UPDATE_PASSWORD: 'admin:user:update:password',
  /** 修改管理员状态 */
  ADMIN_USER_UPDATE_STATUS: 'admin:user:update:status',
  /** 修改管理员角色 */
  ADMIN_USER_UPDATE_ROLE: 'admin:user:update:role',
  /** 上传文件到本地 */
  ADMIN_USER_UPLOAD_FILE: 'admin:user:upload:file',
  /** 查询管理员角色列表 */
  ADMIN_ROLE_FIND_LIST: 'admin:role:find:list',
  /** 创建管理员角色 */
  ADMIN_ROLE_CREATE: 'admin:role:create',
  /** 修改管理员角色信息 */
  ADMIN_ROLE_UPDATE_INFO: 'admin:role:update:info',
  /** 删除管理员角色 */
  ADMIN_ROLE_DELETE: 'admin:role:delete',
  /** 修改管理角色权限码 */
  ADMIN_ROLE_UPDATE_CODE: 'admin:role:update:code',
  /** 查询新闻列表 */
  ADMIN_NEWS_FIND_LIST: 'admin:news:find:list',
  /** 查询新闻详情 */
  ADMIN_NEWS_FIND_DETAIL: 'admin:news:find:detail',
  /** 创建新闻 */
  ADMIN_NEWS_CREATE: 'admin:news:create',
  /** 修改新闻信息 */
  ADMIN_NEWS_UPDATE_INFO: 'admin:news:update:info',
  /** 删除新闻 */
  ADMIN_NEWS_DELETE: 'admin:news:delete',
} as const;
  
export const PERMISSION_MAP = {
  [PERMISSION_CODE.ADMIN_C_USER_FIND_LIST]: '查询C端用户列表',
  [PERMISSION_CODE.ADMIN_C_USER_UPDATE_STATUS]: '修改C端用户状态',
  [PERMISSION_CODE.ADMIN_USER_FIND_LIST]: '查询管理员列表',
  [PERMISSION_CODE.ADMIN_USER_CREATE]: '创建管理员',
  [PERMISSION_CODE.ADMIN_USER_UPDATE_INFO]: '修改管理员基本信息',
  [PERMISSION_CODE.ADMIN_USER_DELETE]: '删除管理员',
  [PERMISSION_CODE.ADMIN_USER_UPDATE_PASSWORD]: '修改管理员密码',
  [PERMISSION_CODE.ADMIN_USER_UPDATE_STATUS]: '修改管理员状态',
  [PERMISSION_CODE.ADMIN_USER_UPDATE_ROLE]: '修改管理员角色',
  [PERMISSION_CODE.ADMIN_USER_UPLOAD_FILE]: '上传文件到本地',
  [PERMISSION_CODE.ADMIN_ROLE_FIND_LIST]: '查询管理员角色列表',
  [PERMISSION_CODE.ADMIN_ROLE_CREATE]: '创建管理员角色',
  [PERMISSION_CODE.ADMIN_ROLE_UPDATE_INFO]: '修改管理员角色信息',
  [PERMISSION_CODE.ADMIN_ROLE_DELETE]: '删除管理员角色',
  [PERMISSION_CODE.ADMIN_ROLE_UPDATE_CODE]: '修改管理角色权限码',
  [PERMISSION_CODE.ADMIN_NEWS_FIND_LIST]: '查询新闻列表',
  [PERMISSION_CODE.ADMIN_NEWS_FIND_DETAIL]: '查询新闻详情',
  [PERMISSION_CODE.ADMIN_NEWS_CREATE]: '创建新闻',
  [PERMISSION_CODE.ADMIN_NEWS_UPDATE_INFO]: '修改新闻信息',
  [PERMISSION_CODE.ADMIN_NEWS_DELETE]: '删除新闻',
} as const;
