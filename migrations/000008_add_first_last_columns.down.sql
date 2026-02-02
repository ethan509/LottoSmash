-- 008_add_first_last_columns.down.sql
-- 첫번째/마지막 번호 위치 통계 컬럼 제거

ALTER TABLE lotto_analysis_stats
DROP COLUMN IF EXISTS first_count,
DROP COLUMN IF EXISTS last_count;
