from topbeat import TestCase


"""
Contains tests for per process statistics.
"""


class Test(TestCase):
    def test_procs(self):
        """
        Checks that the per proc stats are found in the output and
        have the expected types.
        """
        self.render_config_template(
            system_stats=False,
            proc_stats=True,
            filesystem_stats=False,
            proc_patterns=["topbeat.test"]  # monitor itself
        )
        topbeat = self.start_topbeat()
        self.wait_until(lambda: self.output_has(lines=1))
        topbeat.kill_and_wait()

        output = self.read_output()[0]

        print output["proc.name"]
        assert output["proc.name"] == "topbeat.test"
        assert output["proc.state"] == "running"
        assert isinstance(output["proc.cpu.start_time"], basestring)

        for key in [
            "proc.pid",
            "proc.ppid",
            "proc.cpu.user",
            "proc.cpu.system",
            "proc.cpu.total",
            "proc.mem.size",
            "proc.mem.rss",
            "proc.mem.share",
        ]:
            assert type(output[key]) is int

        for key in [
            "proc.cpu.user_p",
            "proc.mem.rss_p",
        ]:
            assert type(output[key]) in [int, float]
