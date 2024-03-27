-- ----------------------------
-- Records of skill
-- ----------------------------
DROP TABLE IF EXISTS "public"."skill";
CREATE TABLE "public"."skill" (
      "id" int8 NOT NULL DEFAULT nextval('skill_id_seq'::regclass),
      "parent_id" int8,
      "zh" varchar(20) COLLATE "pg_catalog"."default",
      "en" varchar(20) COLLATE "pg_catalog"."default",
      "sort" int2 DEFAULT 0,
      "index" int2 DEFAULT 0
)
;
BEGIN;
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (28, 13, 'Rust', 'Rust', 3, 45);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (29, 13, 'C/C++', 'C/C++', 4, 46);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (30, 13, 'PHP', 'PHP', 5, 47);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (31, 13, '.Net', '.Net', 6, 48);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (32, 13, 'SQL', 'SQL', 7, 49);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (33, 13, 'Spring', 'Spring', 8, 64);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (34, 13, 'Laravel', 'Laravel', 9, 65);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (35, 12, 'React', 'React', 3, 66);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (36, 12, 'Vue.js', 'Vue.js', 4, 67);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (37, 14, 'C/C++', 'C/C++', 0, 46);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (1, 0, '开发', 'Development', 0, 1);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (2, 0, '产品', 'Product', 1, 2);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (3, 0, '设计', 'Design', 2, 3);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (4, 0, '测试', 'Testing', 3, 4);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (5, 0, '写作', 'Writing', 4, 5);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (6, 0, '市场营销', 'Marketing', 5, 6);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (7, 0, '调研', 'Research', 6, 7);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (9, 0, '翻译', 'Translation', 8, 8);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (10, 0, '培训', 'Training', 9, 9);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (11, 0, 'Services', 'Services', 10, 10);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (8, 0, '运维', 'Operation', 7, 11);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (12, 1, '前端开发', 'Web Frontend', 0, 20);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (13, 1, '后端开发', 'Backend', 1, 21);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (14, 1, '桌面应用', 'Desktop Apps', 2, 22);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (15, 1, 'Android/iOS', 'Android/iOS', 3, 23);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (16, 1, '自动化/机器人', 'Auto/Bots', 4, 24);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (17, 1, '小程序', 'Mini Program', 5, 25);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (18, 1, '全栈', 'Full Stack', 6, 26);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (19, 1, 'DevOps', 'DevOps', 7, 27);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (20, 1, '数据分析', 'Data Analysts', 8, 28);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (21, 1, '区块链', 'Blockchain', 9, 29);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (22, 1, '人工智能/深度学习', 'AI/ML', 10, 30);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (23, 12, 'JavaScript', 'JavaScript', 0, 40);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (24, 12, 'HTML/CSS', 'HTML/CSS', 1, 51);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (25, 13, 'Python', 'Python', 0, 41);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (26, 13, 'Java', 'Java', 1, 42);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (27, 13, 'Golang', 'Golang', 2, 43);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (38, 14, '.Net', '.Net', 1, 48);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (39, 21, 'Solidity', 'Solidity', 0, 44);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (40, 21, 'Move', 'Move', 1, 50);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (41, 21, 'Hardhat', 'Hardhat', 2, 61);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (42, 21, 'Defi', 'Defi', 3, 62);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (43, 21, 'NFT', 'NFT', 4, 63);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (60, 19, 'Python', 'Python', 1, 41);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (44, 15, 'Java', 'Java', 0, 42);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (45, 16, 'Python', 'Python', 0, 41);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (46, 17, 'HTML/CSS', 'HTML/CSS', 0, 51);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (62, 20, 'SQL', 'SQL', 0, 49);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (63, 20, 'Python', 'Python', 1, 41);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (49, 18, 'Java', 'Java', 0, 42);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (47, 17, 'Vue.js', 'Vue.js', 1, 67);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (48, 17, 'JavaScript', 'JavaScript', 2, 40);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (64, 22, 'Python', 'Python', 0, 41);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (51, 18, 'Golang', 'Golang', 1, 43);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (52, 18, 'C/C++', 'C/C++', 2, 46);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (53, 18, 'Rust', 'Rust', 3, 45);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (54, 18, 'PHP', 'PHP', 4, 47);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (55, 18, '.Net', '.Net', 5, 48);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (50, 18, 'SQL', 'SQL', 6, 49);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (56, 18, 'React', 'React', 7, 66);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (57, 18, 'Vue.js', 'Vue.js', 8, 67);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (58, 18, 'JavaScript', 'JavaScript', 9, 40);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (59, 18, 'HTML/CSS', 'HTML/CSS', 10, 51);
INSERT INTO "public"."skill" ("id", "parent_id", "zh", "en", "sort", "index") VALUES (61, 19, 'Golang', 'Golang', 0, 43);
COMMIT;
-- ----------------------------
-- Records of message_tmpl
-- ----------------------------
DROP TABLE IF EXISTS "public"."message_tmpl";
CREATE TABLE "public"."message_tmpl" (
     "id" int8 NOT NULL DEFAULT nextval('message_tmpl_id_seq'::regclass),
     "issuer" varchar(500) COLLATE "pg_catalog"."default",
     "issuer_zh" varchar(500) COLLATE "pg_catalog"."default",
     "worker" varchar(500) COLLATE "pg_catalog"."default",
     "worker_zh" varchar(500) COLLATE "pg_catalog"."default",
     "status" text COLLATE "pg_catalog"."default",
     "type" int2 DEFAULT 0,
     "disable" bool DEFAULT false,
     "status,unique" varchar(30) COLLATE "pg_catalog"."default"
)
;
BEGIN;
INSERT INTO "public"."message_tmpl" ("id", "issuer", "issuer_zh", "worker", "worker_zh", "status", "type", "disable") VALUES (24, 'The requester has changed the Phase division.<br />To confirm the content of the current Phase division.', '需求方修改了阶段划分，<br />请确认当前阶段划分内容', '', NULL, 'TipWaitIssuerAgree', 1, 'f');
INSERT INTO "public"."message_tmpl" ("id", "issuer", "issuer_zh", "worker", "worker_zh", "status", "type", "disable") VALUES (25, NULL, NULL, 'The project party has modified the phase division.<br />To confirm the content of the current Phase division.', '项目方修改了阶段划分，<br />请确认当前阶段划分内容', 'TipWaitWorkerConfirmStage', 1, 'f');
INSERT INTO "public"."message_tmpl" ("id", "issuer", "issuer_zh", "worker", "worker_zh", "status", "type", "disable") VALUES (1, 'You have successfully published a <a href="/applylist?task_id={{.task_id}}">「{{.title}}」</a> task. <a href="/applylist?task_id={{.task_id}}">view now</a>', '你成功发布了一个任务<a href="/applylist?task_id={{.task_id}}">「{{.title}}」</a>，<a href="/applylist?task_id={{.task_id}}">立即查看</a>', NULL, NULL, 'TaskCreated', 0, 'f');
INSERT INTO "public"."message_tmpl" ("id", "issuer", "issuer_zh", "worker", "worker_zh", "status", "type", "disable") VALUES (3, '「{{.username}}」submitted the phase division of <a href="/order?w=issuer&order_id={{.order_id}}>「{{.title}}」</a>", please confirm. <a href="/order?w=issuer&order_id={{.order_id}}">view now</a>', '「{{.username}}」提交了任务<a href="/order?w=issuer&order_id={{.order_id}}">「{{.title}}」</a>的阶段划分，请前去确认，<a href="/order?w=issuer&order_id={{.order_id}}">立即查看</a>', NULL, NULL, 'WaitIssuerAgree', 0, 'f');
INSERT INTO "public"."message_tmpl" ("id", "issuer", "issuer_zh", "worker", "worker_zh", "status", "type", "disable") VALUES (5, '「{{.username}}」has submitted the {{.stage}} deliverable of the <a href="/order?w=issuer&order_id={{.order_id}}">「{{.title}}」</a> task. <a href="/order?w=issuer&order_id={{.order_id}}">view now</a>', '「{{.username}}」提交了任务<a href="/order?w=issuer&order_id={{.order_id}}">「{{.title}}」</a>的{{.stage}}阶段交付物，<a href="/order?w=issuer&order_id={{.order_id}}">立即查看</a>', NULL, NULL, 'WorkerDelivery', 0, 'f');
INSERT INTO "public"."message_tmpl" ("id", "issuer", "issuer_zh", "worker", "worker_zh", "status", "type", "disable") VALUES (21, '「{{.username}}」agrees to add Phase「{{.stage_name}}」to task <a href="/order?w=issuer&order_id={{.order_id}}">「{{.title}}」</a>. <a href="/order?w=issuer&order_id={{.order_id}}">view now</a>', '「{{.username}}」同意在任务<a href="/order?w=issuer&order_id={{.order_id}}">「{{.title}}」</a>中增加阶段「{{.stage_name}}」，<a href="/order?w=issuer&order_id={{.order_id}}">立即查看</a>', NULL, NULL, 'WaitAppendPayment', 0, 'f');
INSERT INTO "public"."message_tmpl" ("id", "issuer", "issuer_zh", "worker", "worker_zh", "status", "type", "disable") VALUES (4, '「{{.username}}」agrees to modify the task <a href="/order?w=issuer&order_id={{.order_id}}">「{{.title}}」</a>, please confirm. <a href="/order?w=issuer&order_id={{.order_id}}">view now</a>', '「{{.username}}」同意了任务<a href="/order?w=issuer&order_id={{.order_id}}">「{{.title}}」</a>的修改，<a href="/order?w=issuer&order_id={{.order_id}}">立即查看</a>', NULL, NULL, 'WorkerAgreeStage', 0, 'f');
INSERT INTO "public"."message_tmpl" ("id", "issuer", "issuer_zh", "worker", "worker_zh", "status", "type", "disable") VALUES (8, '「{{.username}}」has abrogated the task <a href="/order?w=issuer&order_id={{.order_id}}">「{{.title}}」</a>, you can find new collaborators in the registration list. <a href="/user/projects?w=issuer&bar=tasks">view now</a>', '「{{.username}}」中止了任务<a href="/applylist?task_id={{.order_id}}">「{{.title}}」</a>，你可以在报名列表寻找新的合作者，<a href="/user/projects?w=issuer&bar=tasks">立即查看</a>', '「{{.username}}」abrogated the task <a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a>, you can find other tasks. <a href="/projects">view now</a>', '「{{.username}}」中止了任务<a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a>，你可以寻找其他项目，<a href="/projects">寻找任务</a>', 'OrderAbort', 0, 'f');
INSERT INTO "public"."message_tmpl" ("id", "issuer", "issuer_zh", "worker", "worker_zh", "status", "type", "disable") VALUES (19, '「{{.username}}」has declined to postpone the  {{.stage}} Phase of the <a href="/order?w=issuer&order_id={{.order_id}}">「{{.title}}」</a> task. <a href="/order?w=issuer&order_id={{.order_id}}">view now</a>', '「{{.username}}」拒绝延期任务<a href="/order?w=issuer&order_id={{.order_id}}">「{{.title}}」</a>的{{.stage}}阶段，<a href="/order?w=issuer&order_id={{.order_id}}">立即查看</a>', '「{{.username}}」refused to postpone the {{.stage}} Phase of the task <a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a>. <a href="/order?w=worker&order_id={{.order_id}}">view now</a>', '「{{.username}}」拒绝延期任务<a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a>的{{.stage}}阶段，<a href="/order?w=worker&order_id={{.order_id}}">立即查看</a>', 'DisagreeProlong', 0, 'f');
INSERT INTO "public"."message_tmpl" ("id", "issuer", "issuer_zh", "worker", "worker_zh", "status", "type", "disable") VALUES (2, 'The task <a href="/applylist?task_id={{.task_id}}">「{{.title}}」</a> you published has received a new registration: applicant「{{.username}}」. <a href="/applylist?task_id={{.task_id}}">view now</a>', '你发布的任务<a href="/applylist?task_id={{.task_id}}">「{{.title}}」</a>，收到一个新的报名：报名者-「{{.username}}」，查看<a href="/applylist?task_id={{.task_id}}">报名列表</a>', 'You have successfully signed up for the task <a href="/task/{{.task_id}}">「{{.title}}」</a>. <a href="/task/{{.task_id}}">view now</a>', '你成功报名了任务<a href="/task/{{.task_id}}">「{{.title}}」</a>，<a href="/task/{{.task_id}}">立即查看</a>', 'ApplyFor', 0, 'f');
INSERT INTO "public"."message_tmpl" ("id", "issuer", "issuer_zh", "worker", "worker_zh", "status", "type", "disable") VALUES (9, 'The task <a href="/order?w=issuer&order_id={{.order_id}}">「{{.title}}」</a> is completed, congratulations for getting the NFT. <a href="/myInfo">view now</a>', '任务<a href="/order?w=issuer&order_id={{.order_id}}">「{{.title}}」</a>已完成，恭喜获得NFT，<a href="/myInfo">立即查看</a>', 'The task <a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a> has been completed, congratulations for getting the NFT. <a href="/myInfo">view now</a>', '任务<a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a>已完成，恭喜获得NFT，<a href="/myInfo">立即查看</a>', 'OrderDone', 0, 'f');
INSERT INTO "public"."message_tmpl" ("id", "issuer", "issuer_zh", "worker", "worker_zh", "status", "type", "disable") VALUES (23, NULL, NULL, 'The task <a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a> has a balance to be claimed. <a href="/order?w=worker&order_id={{.order_id}}">view now</a>', '任务<a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a>有待领取余额，<a href="/order?w=worker&order_id={{.order_id}}">立即查看</a>', 'PendingWithdraw', 0, 'f');
INSERT INTO "public"."message_tmpl" ("id", "issuer", "issuer_zh", "worker", "worker_zh", "status", "type", "disable") VALUES (22, NULL, NULL, 'The signature of the task <a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a> has expired, please sign again. <a href="/order?w=worker&order_id={{.order_id}}">view now</a>', '任务<a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a>的签名已失效，请重新签名，<a href="/order?w=worker&order_id={{.order_id}}">立即查看</a>', 'InvalidSign', 0, 'f');
INSERT INTO "public"."message_tmpl" ("id", "issuer", "issuer_zh", "worker", "worker_zh", "status", "type", "disable") VALUES (12, NULL, NULL, '「{{.username}}」paid overall cost of the task <a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a>. <a href="/order?w=worker&order_id={{.order_id}}">view now</a>', '「{{.username}}」支付了任务<a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a>的全部费用，请进行阶段交付，<a href="/order?w=worker&order_id={{.order_id}}">立即查看</a>', 'OrderStarted', 0, 'f');
INSERT INTO "public"."message_tmpl" ("id", "issuer", "issuer_zh", "worker", "worker_zh", "status", "type", "disable") VALUES (10, NULL, NULL, '「{{.username}}」 invites you to Phase the task <a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a>. <a href="/order?w=worker&order_id={{.order_id}}">view now</a>', '「{{.username}}」邀请你对任务<a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a>进行阶段划分，<a href="/order?w=worker&order_id={{.order_id}}">立即查看</a>', 'OrderCreated', 0, 'f');
INSERT INTO "public"."message_tmpl" ("id", "issuer", "issuer_zh", "worker", "worker_zh", "status", "type", "disable") VALUES (11, NULL, NULL, '「{{.username}}」modifies the phase division of the task <a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a>. <a href="/order?w=worker&order_id={{.order_id}}">view now</a>', '「{{.username}}」修改了任务<a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a>的阶段划分，等待你确认，<a href="/order?w=worker&order_id={{.order_id}}">立即查看</a>', 'WaitWorkerConfirmStage', 0, 'f');
INSERT INTO "public"."message_tmpl" ("id", "issuer", "issuer_zh", "worker", "worker_zh", "status", "type", "disable") VALUES (13, NULL, NULL, '「{{.username}}」confirmed the delivery of the task <a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a>. <a href="/order?w=worker&order_id={{.order_id}}">view now</a>', '「{{.username}}」确认了任务<a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a>的交付，<a href="/order?w=worker&order_id={{.order_id}}">立即查看</a>', 'ConfirmOrderStage', 0, 'f');
INSERT INTO "public"."message_tmpl" ("id", "issuer", "issuer_zh", "worker", "worker_zh", "status", "type", "disable") VALUES (18, '「{{.username}}」agrees to add a Phase「{{.stage_name}}」to the task <a href="/order?w=issuer&order_id={{.order_id}}">「{{.title}}」</a>. <a href="/order?w=issuer&order_id={{.order_id}}">view now</a>', '「{{.username}}」同意在任务<a href="/order?w=issuer&order_id={{.order_id}}">「{{.title}}」</a>中增加阶段「{{.stage_name}}」，<a href="/order?w=issuer&order_id={{.order_id}}">立即查看</a>', '「{{.username}}」agrees to add a Phase「{{.stage_name}}」to the task <a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a>. <a href="/order?w=worker&order_id={{.order_id}}">view now</a>', '「{{.username}}」同意在任务<a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a>中增加阶段「{{.stage_name}}」，<a href="/order?w=worker&order_id={{.order_id}}">立即查看</a>', 'AgreeAppend', 0, 'f');
INSERT INTO "public"."message_tmpl" ("id", "issuer", "issuer_zh", "worker", "worker_zh", "status", "type", "disable") VALUES (17, '「{{.username}}」agrees to move the {{.stage}} Phase of the <a href="/order?w=issuer&order_id={{.order_id}}">「{{.title}}」</a> task. <a href="/order?w=issuer&order_id={{.order_id}}">view now</a>', '「{{.username}}」同意延期任务<a href="/order?w=issuer&order_id={{.order_id}}">「{{.title}}」</a>的{{.stage}}阶段，<a href="/order?w=issuer&order_id={{.order_id}}">立即查看</a>', '「{{.username}}」agrees to postpone the {{.stage}} phase of the task <a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a>. <a href="/order?w=worker&order_id={{.order_id}}">view now</a>', '「{{.username}}」同意延期任务<a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a>的{{.stage}}阶段，<a href="/order?w=worker&order_id={{.order_id}}">立即查看</a>', 'AgreeProlong', 0, 'f');
INSERT INTO "public"."message_tmpl" ("id", "issuer", "issuer_zh", "worker", "worker_zh", "status", "type", "disable") VALUES (20, '「{{.username}}」declines to add the Phase「{{.stage_name}}」to the task <a href="/order?w=issuer&order_id={{.order_id}}"> 「{{.title}}」</a>. <a href="/order?w=issuer&order_id={{.order_id}}">view now</a>', '「{{.username}}」拒绝在任务<a href="/order?w=issuer&order_id={{.order_id}}">「{{.title}}」</a>中增加阶段「{{.stage_name}}」，<a href="/order?w=issuer&order_id={{.order_id}}">立即查看</a>', '「{{.username}}」refuses to add Phase「{{.stage_name}}」to task <a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a>. <a href="/order?w=worker&order_id={{.order_id}}">view now</a>', '「{{.username}}」拒绝在任务<a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a>中增加阶段「{{.stage_name}}」，<a href="/order?w=worker&order_id={{.order_id}}">立即查看</a>', 'DisagreeAppend', 0, 'f');
INSERT INTO "public"."message_tmpl" ("id", "issuer", "issuer_zh", "worker", "worker_zh", "status", "type", "disable") VALUES (7, '「{{.username}}」requests to add the Phase「{{.stage_name}}」to the task <a href="/order?w=issuer&order_id={{.order_id}}">「{{.title}}」</a>. <a href="/order?w=issuer&order_id={{.order_id}}">view now</a>', '「{{.username}}」在任务<a href="/order?w=issuer&order_id={{.order_id}}">「{{.title}}」</a>中请求增加「{{.stage_name}}」阶段，<a href="/order?w=issuer&order_id={{.order_id}}">立即查看</a>', '「{{.username}}」requests to add the Phase「{{.stage_name}}」in the task <a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a>. <a href="/order?w=worker&order_id={{.order_id}}">view now</a>', '「{{.username}}」在任务<a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a>中请求增加「{{.stage_name}}」阶段，<a href="/order?w=worker&order_id={{.order_id}}">立即查看</a>', 'WaitAppendAgree', 0, 'f');
INSERT INTO "public"."message_tmpl" ("id", "issuer", "issuer_zh", "worker", "worker_zh", "status", "type", "disable") VALUES (6, '「{{.username}}」requests to extension the {{.stage}} Phase of the task <a href="/order?w=issuer&order_id={{.order_id}}">「{{.title}}」</a>. <a href="/order?w=issuer&order_id={{.order_id}}">view now</a>', '「{{.username}}」请求延期任务<a href="/order?w=issuer&order_id={{.order_id}}">「{{.title}}」</a>的{{.stage}}阶段，请确认，<a href="/order?w=issuer&order_id={{.order_id}}">立即查看</a>', '「{{.username}}」requests to postpone the {{.stage}} Phase of the task <a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a>. <a href="/order?w=worker&order_id={{.order_id}}">view now</a>', '「{{.username}}」请求延期任务<a href="/order?w=worker&order_id={{.order_id}}">「{{.title}}」</a>的{{.stage}}阶段，请确认，<a href="/order?w=worker&order_id={{.order_id}}">立即查看</a>', 'WaitProlongAgree', 0, 'f');
COMMIT;


