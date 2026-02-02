-- 011_add_first_last_prob_columns.sql (down)
-- 첫번째/마지막 번호 위치 확률 컬럼 삭제

ALTER TABLE lotto_analysis_stats
DROP COLUMN IF EXISTS first_prob,
DROP COLUMN IF EXISTS last_prob;
